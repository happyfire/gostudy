package main

import "fmt"

//图graph的key类型是一个字符串，value类型是map[string]bool代表一个字符串集合
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("one", "two")
	fmt.Println(hasEdge("one", "two"))
	fmt.Println(hasEdge("two", "one"))
	fmt.Println(hasEdge("a", "b"))
}
