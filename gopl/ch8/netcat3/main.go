package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		log.Println("receive done", "err:", err)
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)

	if tc, ok := conn.(*net.TCPConn); ok {
		log.Println("shutdown write")
		tc.CloseWrite()
	} else {
		log.Println("close")
		conn.Close()
	}

	<-done // wait for background goroutine to finish
	conn.Close()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	} else {
		log.Println("send done.")
	}
}
