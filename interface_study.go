package main

import (
	"fmt"
	"math"
	"time"
	"io"
	"strings"
	"os"
)

type Abser interface {
	Abs() float64
}

func testInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vector2D{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a Vector2D implements Abser

	a = v // a Vector2D implemenets Abser

	fmt.Println("interface abs:", a.Abs())
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
//func (t T) M() {
//	fmt.Println(t.S)
//}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func testInterface2() {
	var i I = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	var i2 I
	describe(i2)
	//i2.M()

	var anyi interface{}
	describe(anyi)

	anyi = 42
	describe(anyi)

	anyi = "hello"
	describe(anyi)
}

//func describe(i I){
//	fmt.Printf("(%v, %T)\n", i, i)
//}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T! (%v)\n", v, v)
	}
}

func testInterface3() {

	//type assertions
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) //panic
	//fmt.Println(f)

	//type switches
	do(21)
	do("hello")
	do(true)
}

// Stringer interfce (fmt.Stringer)
// type Stringer interface {
//     String() string
// }

type Person struct {
	Name string
	Age  int
}

// Stringer interface method String()
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func testInterfaceStringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	hosts := map[string]IPAddr {
		"loopback" : {127, 0, 0, 1},
		"googleDNS" : {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// error interface

//type error interface {
//	Error() string
//}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrte(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := 1.0
	var last_z = 0.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		if last_z > 0 && (last_z-z) < 0.00000001 {
			break
		}
		last_z = z
	}
	return z, nil
}

func testInterfaceError(){
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrte(2))
	fmt.Println(Sqrte(-2))
}

// io.Reader interface
// func (T) Read(b []byte) (n int, err error)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (n int, err error){
	n, err = r13.r.Read(b)
	if err != nil || n == 0 {
		return n, err
	}

	for i:=0; i<n; i++ {
		var s, e byte
		if b[i] >= 'A' && b[i] <= 'Z' {
			s = 'A'
			e = 'Z'
		} else if b[i] >= 'a' && b[i] <= 'z' {
			s = 'a'
			e = 'z'
		} else {
			continue
		}

		b[i] = (b[i] + 13)
		if b[i] > e {
			b[i] = b[i] - e + s - 1
		}
	}

	return n, err
}

func testInterfaceReader(){
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rs := rot13Reader{s}
	io.Copy(os.Stdout, &rs)
	fmt.Println("\n")
}