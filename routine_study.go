package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func testRoutine() {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		//fmt.Println(sum)
	}
	c <- sum // send sum to channel c
}

func testBufferdChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3

	fmt.Println("buffered", <-ch)
	fmt.Println("buffered", <-ch)
}

func testChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	//buffer channel
	go testBufferdChannel()

	testCloseChannel()

	testChannelSelect()

	time.Sleep(time.Second)
}

func fibonacci_c(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func testCloseChannel() {
	c := make(chan int, 10)
	go fibonacci_c(cap(c), c)
	//如果channel没有close，使用range就会阻塞
	for i := range c {
		fmt.Println(i)
	}
	//close channel之后再读取channel，会得到空值
	fmt.Println("read channel after close:", <-c)
}

func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			// no blocking here
		}
	}
}

func testChannelSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci_select(c, quit)
}
