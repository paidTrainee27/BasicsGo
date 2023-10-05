package main

import "fmt"

//FIFO
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(v int) {
	q.items = append(q.items, v)

}
func (q *Queue) Dequeue() {
	q.items = q.items[1:]

}
func (q Queue) PrintOut() {
	fmt.Println(q.items)
}

func main() {
	q := new(Queue)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(7)
	q.Enqueue(1)
	q.PrintOut()
	q.Dequeue()
	q.PrintOut()

}
