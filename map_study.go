package main

import (
	"fmt"
	"strings"
)

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

var m2 = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.43234, -74.2333,
	},
	"Google": Vertex2{
		37.43, -122.3432,
	},
}

var m3 = map[string]Vertex2{
	"Bell Labs": {40.34234, -74.234234},
	"Google":    {37.32423, -122.3232},
}

func testMap() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m2)

	fmt.Println(m3)

	m4 := make(map[string]int)
	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	v, ok := m4["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println(wordCount("I am learning Go! Go Go Go!"))

}

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)

	for i := 0; i < len(words); i++ {
		result[words[i]]++
	}

	return result
}
