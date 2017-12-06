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
	fmt.Println("====main====")
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

	//test method value
	{
		p := geometry.Point{1, 2}
		q := geometry.Point{4, 6}

		distanceFromP := p.Distance        // method value
		fmt.Println(distanceFromP(q))      //5
		var origin geometry.Point          // {0 0}
		fmt.Println(distanceFromP(origin)) // 23606797749979, sqrt(5)

		scaleP := p.ScaleBy // method value
		scaleP(2)           // p becomes (2,4)
		scaleP(3)           // then (6,12)
		scaleP(10)          // then (60, 120)
		fmt.Println(p)
	}

	//test method expression
	{
		p := geometry.Point{1, 2}
		q := geometry.Point{4, 6}

		distance := geometry.Point.Distance // method expression
		fmt.Println(distance(p, q))         // 5
		fmt.Printf("%T\n", distance)        // func(geometry.Point, geometry.Point) float64

		scale := (*geometry.Point).ScaleBy
		scale(&p, 2)
		fmt.Println(p)            // {2 4}
		fmt.Printf("%T\n", scale) //func(*geometry.Point, float64)
	}
}
