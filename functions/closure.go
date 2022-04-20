package main

import "fmt"

type cls1 func() int

func callClosure() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

func intSeq() cls1 {
	i := 0
	return func() int {
		i++
		return i
	}
}
