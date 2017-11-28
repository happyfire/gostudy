package main

import "fmt"

func main() {

	months := [...]string{1: "January", 2: "February", 3: "March",
		4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]     //第二季度
	summer := months[6:9] //夏季，6,7,8三个月
	fmt.Println(Q2)       //[April May June]
	fmt.Println(summer)   //[June July August]

	//两个切片底层是同一数组，切片可以重叠
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	//切片只能和nil比较
	// fmt.Println(summer == Q2) //invalid operation: summer == Q2 (slice can only be compared to nil)

	//超出容量panic
	//fmt.Println(summer[:20]) //panic: runtime error: slice bounds out of range

	endlessSummer := summer[:5] //在容量限制内,新的切片长度变大
	fmt.Println(endlessSummer)  //[June July August September October]

	//{
	//	str := "hello"
	//	a := str[0:2]
	//	fmt.Printf("%p %v\n", &str, str)
	//	fmt.Printf("%p %v\n", &a, a)
	//}
	//{
	//	buf := []byte{1, 2, 3}
	//	fmt.Printf("%p %v\n", &buf, buf)
	//	a := buf[0:2]
	//	fmt.Printf("%p %v\n", &a, a)
	//}

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) //['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']

	{
		//stack
		fmt.Println("====stack===")
		stack := []int{}
		stack = append(stack, 1) // push
		fmt.Println(stack)
		stack = append(stack, 2) // push
		fmt.Println(stack)
		top := stack[len(stack)-1] //top of stack
		fmt.Println(stack, top)
		stack = stack[:len(stack)-1] //pop
		fmt.Println(stack)
	}

	{
		s := []int{5, 6, 7, 8, 9}
		fmt.Println(remove(s, 2)) //[5 6 8 9]
	}

	{
		var sli []int
		fmt.Println(sli, sli == nil, len(sli)) //nil slice的长度也是0
	}

}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
