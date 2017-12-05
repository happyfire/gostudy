package main

import (
	"fmt"
	"sort"
)

func main() {

	args := make(map[string]int) // mapping from strings to ints
	args["port"] = 8080
	fmt.Println(args)

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages)
	ages["alice"] = 32
	fmt.Println(ages["alice"])
	delete(ages, "alice") // remove element ages["alice"]
	fmt.Println(ages)
	ages["bob"]++
	fmt.Println(ages["bob"])
	ages["tom"] = 0
	if age, ok := ages["jim"]; !ok {
		fmt.Println("jim is not in this map", age)
	}

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//map的遍历顺序是随机的，想获得固定的顺序，可使用一个slice存放key，并对slice进行排序
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	fmt.Println(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	//map类型的零值是nil
	{
		var m map[string]int
		fmt.Println(m["a"], m == nil, len(m) == 0) //0, true, true
		//m["a"] = 10 //panic: assignment to entry in nil map
	}

	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}

		//错误写法：（会导致上面的测试返回true)
		//if xv != y[k] {
		//	return false
		//}
	}
	return true
}
