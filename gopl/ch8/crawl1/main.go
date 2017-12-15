package main

import (
	"fmt"
	"github.com/happyfire/gostudy/gopl/ch5/links"
	"log"
)

// 不限制并发数的并发爬虫，按书上的说法应该会让进程的fd很快耗尽
// 并且这个程序不会结束，因为worklist这个channel不会关闭

// 实际我跑的时候没有看到 socketL too many open files这类错误，但是有很多tls timeout, io timeout
// 但是把公司网络搞得不能上网了

func main() {
	worklist := make(chan []string)

	go func() { worklist <- []string{"http://gopl.io/"} }()

	// crawl the web concurrently
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
