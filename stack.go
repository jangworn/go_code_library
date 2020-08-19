package main

import (
	"fmt"
)

// Push 栈添加元素
func (s *Stack) Push(element int) {
	s.Data = append(s.Data, element)
}

// Pop 栈取出元素
func (s *Stack) Pop() {
	len := len(s.Data)
	s.Data = append(s.Data[:len-1], s.Data[len:]...)
}

// Length 返回栈长度
func (s *Stack) Length() int {
	len := len(s.Data)
	return len
}

// Clear 清空栈
func (s *Stack) Clear() {
	s.Data = []int{}
}

// Peek 返回第一个元素
func (s *Stack) Peek() int {
	return s.Data[0]
}

// Stack 栈
type Stack struct {
	Data []int
}

func main() {
	var s Stack
	i := 1
	n := 5
	for {
		if l := s.Length(); l > n {
			break
		}
		s.Push(i)
		i++
	}
	fmt.Println(s.Data)
	s.Pop()
	fmt.Println(s.Data)
	s.Pop()
	fmt.Println(s.Data)
	fmt.Println(s.Peek())
	s.Clear()
	fmt.Println(s.Data)
}
