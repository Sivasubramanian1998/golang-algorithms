package main

import "fmt"

//stack represents stack that holds slice
type Stack struct {
	items []int
}

// Push will add the value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove the value at the end
// and returns the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	remove := s.items[l]
	s.items = s.items[:l]
	return remove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(1000)
	myStack.Push(2000)
	myStack.Push(3000)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
