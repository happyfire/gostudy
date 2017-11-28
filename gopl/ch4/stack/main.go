package main

import "fmt"
import "errors"

type StackInt struct {
	slice []int
}

func (s *StackInt) Push(v int) {
	s.slice = append(s.slice, v)
}

func (s StackInt) Size() int {
	return len(s.slice)
}

func (s StackInt) Top() (int, error) {
	if s.Size() > 0 {
		return s.slice[s.Size()-1], nil
	}
	return 0, errors.New("stack is empty")
}

func (s *StackInt) Pop() (int, error) {
	top, err := s.Top()
	if err == nil {
		s.slice = s.slice[:s.Size()-1]
		return top, nil
	}
	return top, err
}

func main() {
	s := StackInt{}
	fmt.Println(s, s.Size())
	s.Push(1)
	fmt.Println(s, s.Size())
	s.Push(2)
	fmt.Println(s, s.Size())
	s.Push(3)
	fmt.Println(s, s.Size())
	fmt.Println(s.Top())
	fmt.Println(s, s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s, s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s, s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s, s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s, s.Size())
}
