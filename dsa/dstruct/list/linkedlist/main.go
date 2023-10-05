package main

import "fmt"

type NodeInt struct {
	Data     int
	NextNode *NodeInt
}
type LinqListInt struct {
	Head *NodeInt
	Size int
}

func (l *LinqListInt) Contains(value int) bool {
	curr := l.Head
	nxt := curr.NextNode
	if curr.Data == value {
		return true
	}

	for nxt != nil {
		if nxt.Data == value {
			return true
		}
		tempNxt := nxt
		nxt = nxt.NextNode
		curr = tempNxt
	}
	return false
}

func (l *LinqListInt) Remove(value int) {
	curr := l.Head
	nxt := curr.NextNode
	if curr.Data == value {
		l.Head = nxt
		return
	}

	for nxt != nil {
		if nxt.Data == value {
			curr.NextNode = nxt.NextNode
			l.Size--
			return
		}
		tempNxt := nxt
		nxt = nxt.NextNode
		curr = tempNxt
	}
}

func (l *LinqListInt) Empty() {
	l.Head = nil
	l.Size = 0
}

func (l *LinqListInt) Reverse() {
	h := l.Head
	if h == nil {
		fmt.Println("Empty list")
		return
	}
	nxt := l.Head.NextNode
	l1 := LinqListInt{}
	l1.Add(h.Data)
	for nxt != nil {
		l1.Add(nxt.Data)
		nxt = nxt.NextNode
	}
	l.Head = l1.Head
}

func (l *LinqListInt) Add(value int) {
	//SWAP basically
	h := l.Head
	n := new(NodeInt)
	n.Data = value
	l.Head = n
	n.NextNode = h
	l.Size++
}

func (l LinqListInt) SysOut() {
	h := l.Head
	if h == nil {
		fmt.Println("Empty list")
		return
	}
	nxt := l.Head.NextNode
	fmt.Print(h.Data)
	for nxt != nil {
		fmt.Print("->")
		fmt.Print(nxt.Data)
		nxt = nxt.NextNode
	}
	fmt.Println("")
}

func main() {
	l := new(LinqListInt)
	l.Add(2)
	l.Add(4)
	l.Add(1)
	l.Add(20)
	l.Add(10)
	l.Add(24)
	l.Add(28)
	// l.Reverse()
	l.Remove(288)
	// l.Empty()
	l.SysOut()
	// fmt.Println(l.Contains(20))
	// fmt.Println("Hey DeV!")
}
