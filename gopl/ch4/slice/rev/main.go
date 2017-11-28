package main

import "fmt"

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse_bytes(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	s := []byte{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse_bytes(s[:2])
	fmt.Println(s) //[1 0 2 3 4 5]
	reverse_bytes(s[2:])
	fmt.Println(s) //[1 0 5 4 3 2]
	reverse_bytes(s)
	fmt.Println(s) //[2 3 4 5 0 1]
}
