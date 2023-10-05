package main

import "fmt"

type BinaryST struct {
	Value        int
	LeftSubtree  *BinaryST
	RightSubtree *BinaryST
}

//Add
func (b *BinaryST) Add(v int) {
	if v < b.Value {
		if b.LeftSubtree == nil {
			b1 := new(BinaryST)
			b1.Value = v
			b.LeftSubtree = b1
		} else {
			b.LeftSubtree.Add(v)
		}

	}
	if v > b.Value {
		if b.RightSubtree == nil {
			b1 := new(BinaryST)
			b1.Value = v
			b.RightSubtree = b1
		} else {
			b.RightSubtree.Add(v)
		}
	}
}

//Search
func (b *BinaryST) Search(v int) bool {
	if b.Value == v {
		return true
	}

	if v < b.Value && b.LeftSubtree != nil {
		return b.LeftSubtree.Search(v)
	}

	if v > b.Value && b.RightSubtree != nil {
		return b.RightSubtree.Search(v)
	}

	return false
}

//Delete

func main() {
	b := BinaryST{Value: 30}
	b.Add(10)
	b.Add(1)
	b.Add(34)
	b.Add(98)
	b.Add(25)
	fmt.Println(b.Search(14))
}
