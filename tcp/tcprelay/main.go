package main

import (
	"io"
	"log"
	"net"
)

func panicOnError(err error) {
	if err != nil {
		log.Println("panic on error:", err)
		panic(err)
	}
}

func relay(serverconn net.Conn) {
	defer serverconn.Close()
	log.Printf("relay: listen on %v, new client from %v\n",
		serverconn.LocalAddr().String(), serverconn.RemoteAddr().String())

	clientconn, err := net.Dial("tcp", "localhost:3000")
	panicOnError(err)
	defer clientconn.Close()
	log.Printf("relay: connected to %v, from %v\n",
		clientconn.RemoteAddr().String(), clientconn.LocalAddr().String())

	done := make(chan struct{})
	go func() {
		io.Copy(clientconn, serverconn)
		log.Printf("relay: done copying serverconn to clientconn\n")
		if tcpconn, ok := clientconn.(*net.TCPConn); ok {
			tcpconn.CloseWrite()
		} else {
			clientconn.Close()
		}

		done <- struct{}{}
	}()

	io.Copy(serverconn, clientconn)
	log.Printf("relay: done copying clientconn to serverconn\n")
	if tcpconn, ok := serverconn.(*net.TCPConn); ok {
		tcpconn.CloseWrite()
	} else {
		serverconn.Close()
	}
	<-done
}

func main() {
	listener, err := net.Listen("tcp", ":2000")
	panicOnError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go relay(conn)
	}
}
