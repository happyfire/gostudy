package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", x and not y {5}
	fmt.Printf("%08b\n", ^x)   // "11011101"

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	xx := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", xx) // "3735928559 deadbeef 0xdeadbeef 0XDEADBEEF"

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    //"true"!

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) //"false false false"

	s := "hello, world"
	fmt.Println(len(s)) // 12
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])            //hello
	fmt.Println(s[:5])             //hello
	fmt.Println(s[7:])             //world
	fmt.Println(s[:])              //hello, world
	fmt.Println("goodbye" + s[5:]) //goodbye, world

	const GoUsage = `Go is a tool for managing Go Source code.

	Usage:
		go command [arguments]`

	fmt.Println(GoUsage)

	s = "Hello, 世界"
	fmt.Println(len(s))                    // 13
	fmt.Println(utf8.RuneCountInString(s)) //9

	n := 0
	for range s {
		n++
	}
	fmt.Println(n)

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%x\n", i, r, r)
	}

	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))

	fmt.Println(string(65)) //A
	fmt.Println(string(0x4e16))
	fmt.Println(string(1234567)) //无效的unicode码点
}
