package main

import (
	"bytes"
	"fmt"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer

	m := len(s) % 3

	if m > 0 {
		buf.WriteString(s[0:m])
	}

	for i := m; i < len(s); i += 3 {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234567"))

	fmt.Println(comma2("1234567"))

	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s, b, s2)
}
