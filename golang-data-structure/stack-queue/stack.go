package main

import "fmt"

type Stack struct {
	items []int
}

// Push : menambahkan value pada bagian atas stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop: menghapus value pada bagian atas stack
func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
