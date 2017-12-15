package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//这个版本使用sync.WaitGroup等待所有echo goroutine执行完成后再close connection
//这样客户端就能收到所有的回文（当然客户端必须是closeWrite而不是Close，netcat3就是这样）

func handleConn(c net.Conn) {
	var wg sync.WaitGroup // number of echo goroutines

	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
	}
	// Note: ignoring potential errors from input.Err()

	//这儿因为Wait之后没有可能引发死锁的channel操作，所以可以直接放在主goroutine里面wait，否则要另起一个goroutine
	wg.Wait()
	c.Close()

}
