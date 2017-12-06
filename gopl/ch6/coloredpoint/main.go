package main

import "image/color"
import (
	"fmt"
	"github.com/happyfire/gostudy/gopl/ch6/geometry"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{geometry.Point{1, 1}, red}
	var q = ColoredPoint{geometry.Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) //5
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) //10

}

func init() {
	type ColoredPoint struct {
		*geometry.Point
		Color color.RGBA
	}

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	p := ColoredPoint{&geometry.Point{1, 1}, red}
	q := ColoredPoint{&geometry.Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) //5
	q.Point = p.Point                 // p and q now share the same point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // {2 2} {2 2}
}
