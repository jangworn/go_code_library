package main

import "fmt"

func (q *Queue) Push(element int) {
	q.Data = append(q.Data, element)
}

func (q *Queue) Pop() {
	q.Data = append(q.Data[:0], q.Data[1:]...)
}

// Length 返回队列长度
func (q *Queue) Length() int {
	len := len(q.Data)
	return len
}

// Clear 清空队列
func (q *Queue) Clear() {
	q.Data = []int{}
}

// Peek 返回第一个元素
func (q *Queue) Peek() int {
	return q.Data[0]
}

// Queue 队列
type Queue struct {
	Data []int
}

func main() {
	var q Queue
	i := 1
	n := 5
	for {
		if l := q.Length(); l > n {
			break
		}
		q.Push(i)
		i++
	}
	fmt.Println(q.Data)
	q.Pop()
	fmt.Println(q.Data)
	q.Pop()
	fmt.Println(q.Data)
	fmt.Println(q.Peek())
	q.Clear()
	fmt.Println(q.Data)
}
