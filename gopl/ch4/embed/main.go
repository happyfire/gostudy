package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point  //匿名成员
	Radius int
}

type Wheel struct {
	Circle //匿名成员
	Spokes int
}

func main() {

	{
		//结构体字面值没有简短表示匿名成员的语法
		//w := Wheel{8, 8, 5, 20} //编译错误
		//w := Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} //unknown field
	}

	{
		w := Wheel{Circle{Point{8, 8}, 5}, 20}
		fmt.Printf("%#v\n", w)
	}

	{
		w := Wheel{
			Circle: Circle{
				Point:  Point{X: 8, Y: 8},
				Radius: 5, //结尾的逗号是必须的
			},
			Spokes: 20, //结尾的逗号是必须的
		}
		w.X = 42
		fmt.Printf("%#v\n", w)
	}
}
