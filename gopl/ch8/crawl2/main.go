package main

import (
	"fmt"
	"github.com/happyfire/gostudy/gopl/ch5/links"
	"log"
)

// 在worklist为空或者没有crawl的goroutine在运行时退出主循环

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	//n表示worklist中需要处理的list数量
	n++ //启动时先加一，因为将会从下面的goroutine加入一个list到worklist，这保证下面的循环不直接退出，会等待从worklist取出

	go func() { worklist <- []string{"http://gopl.io/"} }()

	// crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++ //每当将有一个list被加入worklist时，增加n
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// 使用带缓冲的channel实现限制并发数
// buffered channel被用作信号量

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}
