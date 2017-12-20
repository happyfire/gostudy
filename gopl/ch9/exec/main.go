package main

import (
	"fmt"
	"time"
)

func ping(in <-chan string, out chan<- string) {
	for {
		out <- "ping"
		fmt.Println(<-in)
		time.Sleep(time.Second)
	}
}

func pong(in <-chan string, out chan<- string) {
	for {
		fmt.Println(<-in)
		out <- "pong"
		time.Sleep(time.Second)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go ping(ch1, ch2)
	go pong(ch2, ch1)

	//go func() {
	//	for {
	//		ch1 <- "ping"
	//		fmt.Println(<-ch2)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		fmt.Println(<-ch1)
	//		ch2 <- "pong"
	//	}
	//}()

	select {}
}
