package main

import (
	"fmt"
	"github.com/happyfire/gostudy/foo"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func add(x, y int) int {
	return x + y
}

func test(x, y int, z int) int {
	return x * y * z
}

func swap(x, y string) (string, string, int) {
	return y, x, 3
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	defer fmt.Println("defer message")
	for i := 0; i < 3; i++ {
		defer fmt.Println("defer", i)
	}

	var c, python, java = true, false, "no!"

	var i, j = 1, "j=2"
	var x, y int = 2, 3

	fmt.Println("My favorite number is", rand.Intn(10))
	foo.PrintPi()
	fmt.Println(test(2, 3, 4))
	fmt.Println(add(42, 13))
	a, b, _c := swap("hello", "world")
	fmt.Println(a, b, _c)
	fmt.Println(split(17))

	k := 3
	w := true
	fmt.Println(c, python, java)
	fmt.Println(i, j, x, y, k, w)

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	var s string
	var bl bool
	fmt.Printf("%q %v\n", s, bl)

	var ff float64 = math.Sqrt(float64(x*x + y*y))
	var zz uint = uint(ff)
	fmt.Println(ff, zz)

	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	testLoop()

	testif()

	testSqrt()

	testSwitch()

	testPointer()

	testStruct()

	testArray()

	testSlice()

	testForRange()

	testMap()

	testFuncValues()

	testFuncClosures()

	testFibonacciClosure()

	testMethod()

	testInterface()
	testInterface2()
	testInterface3()
	testInterfaceStringer()
	testInterfaceError()
	testInterfaceReader()

	testRoutine()
	testChannel()

	testMutex()

	testWebCrawler()
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func testLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum =", sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum =", sum)

}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func testif() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func Sqrt(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}

	z := 1.0
	var last_z = 0.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(i, "->", z, last_z)
		if last_z > 0 && (last_z-z) < 0.00000001 {
			break
		}
		last_z = z
	}
	return z
}

func testSqrt() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
	fmt.Println(Sqrt(-2), math.Sqrt(-2))
}

func testSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func testPointer() {
	i, j := 42, 2701

	p := &i         //point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         //set i through the pointer
	fmt.Println(i)

	var q *int
	q = &j
	*q = *q / 37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func testStruct() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	//结构体字面量
	var (
		v1 = Vertex{1, 2}  //类型为Vertex
		v2 = Vertex{X: 1}  // Y : 0被省略
		v3 = Vertex{}      // X : 0 和 Y : 0
		pv = &Vertex{1, 2} //类型为 *Vertex
	)

	fmt.Println(v1, pv, v2, v3)

}

func testArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(len(a))
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func testSlice() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	//slice切片
	fmt.Println("s[1:2] ==", s[1:2])
	//省略下标代表从0开始
	fmt.Println("s[1:4] ==", s[:3])
	//省略上标代表到len(s)结束
	fmt.Println("s[4:] ==", s[4:])

	//构造slice
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	//nil slices
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("z is nil")
	}
	fmt.Printf("%p\n", &z)
	z = append(z, 111)
	fmt.Println(z)

	//append slice
	var e []int
	printSlice("e", e)
	// append works on nil slices
	e = append(e, 0)
	printSlice("e", e)
	// the slice grows as needed
	e = append(e, 1)
	printSlice("e", e)
	// we can add more than one elemnt at a time
	e = append(e, 2, 3, 4)
	printSlice("e", e)

	//slices-of-slice

	//create a tic-tac-board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func testForRange() {
	//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//for i, v := range pow {
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}

	pow := make([]int, 10)
	for i := range pow { //如果只需要索引值，去掉", value"部分即可
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow { //可以通过赋值给 _ 来忽略索引和值
		fmt.Printf("%d\n", value)
	}

}

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			s[i][j] = uint8(i * j)
		}
	}
	return s
}
