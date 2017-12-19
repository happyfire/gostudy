package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//这个例子演示了如何终止执行中的大量且不知道数量的goroutines
// 1. 通过一个channel done传递终止信号，这个channel不写入，只是在发出信号时关闭它
// 2. 会在多个地方使用select检查done，如果done被关闭，则select到done。
//    首先，新生成的walkDir goroutine开始的时候检查done，如果已经退出则goroutine直接返回
//	  在dirents中检查到done也直接返回,这样避免从阻塞恢复的dirents产生新的goroutine
//    最后，在主goroutine中，检查到done需要排空fileSizes，避免walkDir goroutine阻塞在向fileSizes写入，让这些goroutine都执行完毕

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) //read a single byte
		close(done)
	}()

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish
			for range fileSizes {
				// Do nothing.
			}
			fmt.Println("aborted.")
			return
			//panic("abort")
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20) // concurrency-limiting counting semaphore

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: //acquire token
	case <-done:
		return nil //cancelled
	}
	defer func() { <-sema }() // release token

	// read directory
	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}

func init() {
	tc := make(chan int)
	go func() {
		tc <- 1
		close(tc)
	}()

	fmt.Println(<-tc)
	fmt.Println(<-tc)
	fmt.Println(<-tc)

	for _, _ = range []string(nil) {
		fmt.Println("in loop")
	}

	fmt.Println("ok")

}
