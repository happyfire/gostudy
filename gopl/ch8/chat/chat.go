package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages 来自所有客户端的消息
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// broadcast incoming message to all clients' outgoing message channels
			for cli := range clients {
				cli <- msg //这儿的cli就是handleConn里面的ch
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages 用于将消息发送到tcp连接上，服务器直接发送的和广播的都发到ch上，ch再发到tcp
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch //ch发送到entering，让广播器添加到map中

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text() //从tcp发过来的客户端的消息，写入到messages中
	}
	// NOTE: ignoring potential errors from input.Err()

	//tcp断开，发送立刻消息，且将ch从广播器map中移除
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//func init() {
//	ch := make(chan int)
//
//	go func() { fmt.Println(<-ch) }() //先启动一个goroutine读取ch,然后这个goroutine会阻塞等待
//	ch <- 1                           //再向ch写会立刻执行
//
//	//上面两部如果倒过来就会死锁
//}
