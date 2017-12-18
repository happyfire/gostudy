package main

import (
	"fmt"
	"github.com/happyfire/gostudy/gopl/ch5/links"
	"log"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // 去重的URL (de-duplicated URLs)

	go func() { worklist <- []string{"http://gopl.io/"} }()

	// create 20 crawler goroutines to fetch each unseen link.
	//创建20个工作goroutine，每个goroutine从unseenLinks里面取出Link,通过crawl爬出新的Links，再添加到worklist中
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				//从foundLinks添加到worklist使用一个单独的goroutine避免死锁,因为可能主goroutine还没来得及从worklist中取
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	// 主goroutine将worklist里面的link去重并发送到unseenLinks中，供工作goroutine获取
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
