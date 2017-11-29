package main

import (
	"fmt"
	"github.com/happyfire/gostudy/gopl/ch4/github"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"repo:golang/go", "is:open", "json", "decoder"}
	}

	result, err := github.SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
