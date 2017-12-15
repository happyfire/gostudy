package main

import (
	"fmt"
	"sync"
	"time"
)

//这是一个测试版本，没有去执行缩略图，只是为了看wg的使用
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup //number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			sizes <- 1
		}(f)
	}

	// closer
	//放在一个goroutine里面是因为要让主goroutine下面的range sizes有机会执行
	//因为sizes是一个无缓冲的channel，必须靠range sizes里面的读取去让worker goroutine里面的sizes <- 1解除阻塞，从而可以让wg.Done执行
	//如果wg.Wait放主goroutine里面，则没有一个worker goroutine能结束，从而没有一个wg.Done被执行，因为wg.Wait会永远阻塞，死锁！
	go func() {
		wg.Wait()
		close(sizes)
	}()

	//in main goroutine
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	fns := make(chan string, 3)
	for i := 0; i < 3; i++ {
		fns <- fmt.Sprintf("%d", i)
	}
	close(fns)

	fmt.Println(makeThumbnails6(fns))
}
