package main

import (
	"fmt"
	"net/http"
)

func request(url string) string {
	resp, _ := http.Get(url)
	if resp != nil {
		return url + " returns " + resp.Status
	} else {
		return url + " error"
	}
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("http://www.gopl.io/") }()
	go func() { responses <- request("http://www.gopl.io/reviews.html") }()
	go func() { responses <- request("http://www.gopl.io/errata.html") }()
	return <-responses
}

func main() {
	fmt.Println(mirroredQuery())
}
