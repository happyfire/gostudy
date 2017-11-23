package main

import (
	"fmt"
	"math"
)

type Vector2D struct {
	X, Y float64
}

func (v Vector2D) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2D) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func Abs(v Vector2D) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vector2D, f float64) {
	v.X *= f
	v.Y *= f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func testMethod() {
	v := Vector2D{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	v2 := Vector2D{3, 4}
	Scale(&v2, 10)
	fmt.Println(Abs(v2))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
