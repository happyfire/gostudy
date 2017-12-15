package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// Note: ignoring potential errors from input.Err()
	time.Sleep(3 * time.Second) //注意：这儿等待3秒是为了等待所有的echo完成的简单处理,应该有更科学的处理方法
	c.Close()

	//注：这个例子是书上的reverb2，在echo调用前加了go，因此netcat3只是closeWrite这儿也会因为echo是goruntine而立刻执行close，
	// 导致netcat3不能收到所有的回应，所以这儿临时改了下增加了一个sleep。书上说reverb2处理比较困难确实如此。如果是reverb1，因为echo不是gorunine
	// 会顺序发送完毕然后close

	//注：在reverb3中，使用sync.WaitGroup解决这个问题
}
