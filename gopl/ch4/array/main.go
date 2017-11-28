package main

import (
	"fmt"
)

func main() {

	//声明数组
	var a [3]int //元素初始化为默认零值
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1]) // len 返回元素个数

	//使用数组字面值语法初始化数组
	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q)
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) //未初始化的元素默认零值

	//使用...根据初始化值的个数计算数组长度
	qq := [...]int{1, 2, 3}
	fmt.Printf("%T\n", qq) // [3]int

	//使用索引值对列表的方式初始化
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "欧", GBP: "磅", RMB: "元"} //索引顺序无关紧要
	fmt.Println(RMB, symbol[RMB])

	//定义10个元素的数组rr，最后一个元素被初始化为-1，其他为0
	rr := [...]int{9: -1} //索引可省略，未指定的元素用零值初始化
	fmt.Println(rr)

	//数组比较必须类型相同，且元素可比较
	//直接用==比较，所有元素相等数组才相等
	{
		a := [2]int{1, 2}
		b := [...]int{1, 2}
		c := [2]int{1, 3}
		fmt.Println(a == b, a == c, b == c) //true, false, false
		//d := [3]int{1, 2}
		//fmt.Println(a == d) //invalid operation: a == d (mismatched types [2]int and [3]int)
	}

	var arr [32]byte = [32]byte{31: 1}
	fmt.Printf("%p %v\n", &arr, arr)
	zero(&arr)
	fmt.Printf("%p %v\n", &arr, arr)
}

//go语言如果使用数组作为参数，会拷贝整个数组，因为要使用指针
func zero(ptr *[32]byte) {
	*ptr = [32]byte{} //数组复制是值复制,这儿使用新生成的零时数组去复制值给传入的数组
}
