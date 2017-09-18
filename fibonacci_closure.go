package main

import "fmt"

func fibonacci() func() int {

	s1, s2 := 0, 1

	return func() int {
		r := s1
		s1, s2 = s2, (s1 + s2)
		return r
	}
}

func testFibonacciClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
