package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func testMap() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
}
