package main

import "fmt"

//LIFO
type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)

}
func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]

}
func (s Stack) PrintOut() {
	fmt.Println(s.items)
}

func main() {
	s := new(Stack)
	s.Push(3)
	s.Push(4)
	s.Push(7)
	s.Push(1)
	s.PrintOut()
	s.Pop()
	s.PrintOut()

}
