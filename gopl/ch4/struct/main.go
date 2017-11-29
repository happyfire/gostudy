package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

var AllEmployees = map[int]*Employee{} //因为map存放的元素不能取值，所以这儿只能存放Employee的指针

func EmployeeByID(id int) *Employee {
	return AllEmployees[id]
}

type Point struct{ X, Y int }

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

type Circle struct {
	Point  //匿名成员
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {

	var dilbert Employee
	dilbert.ID = 1
	dilbert.Name = "dilbert"
	dilbert.Salary = 5000
	position := &dilbert.Position //对成员取地址，然后通过指针访问
	*position = "Senior " + *position

	boss := Employee{ID: 2, Name: "boss", Position: "CEO"}

	dilbert.ManagerID = boss.ID

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"    //点操作符可以和指向结构体的指针一起工作
	(*employeeOfTheMonth).Position += " (proactive team player)" //和上面一样效果

	AllEmployees[dilbert.ID] = &dilbert
	AllEmployees[boss.ID] = &boss

	EmployeeByID(dilbert.ID).Salary = 0
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)

	{
		pt := Point{1, 2}
		fmt.Printf("%p, %v\n", &pt, pt)
		pt2 := Scale(pt, 2)
		fmt.Printf("%p, %v\n", &pt2, pt2)

		pp := &Point{1, 2}
		pp2 := new(Point)
		*pp2 = Point{1, 2}
		fmt.Println(pp, pp2, *pp, *pp2)

		fmt.Println(pp == pp2, *pp == *pp2)
	}

	{
		//匿名成员
		var w Wheel
		w.X = 8
		w.Circle.Point.Y = 8
		w.Radius = 5
		w.Spokes = 20
		fmt.Println(w)
	}
}
