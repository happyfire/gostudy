package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		time.Sleep(time.Second * 10) //等待10秒才read
		// 等待过程中，客户端发送完数据已经退出了，但是等待结束后，服务器还是能收到数据以及EOF
		// 说明go底层已经将数据接收下来了，只是在等待goroutine执行

		// read from the connection
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}

		go handleConn(c)
	}
}
