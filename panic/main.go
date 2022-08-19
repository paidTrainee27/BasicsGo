package main

import (
	"fmt"
	"sync"
)

func main() {
	PanicWithGoroutine()
}

func handleBydefer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occured", err)
		}
	}()
	dividByZero(54, 0)
}

func dividByZero(a, b int) int {
	return a / b
}

func handlePanic(callerName string) {
	r := recover()
	if r != nil {
		fmt.Printf("recovered in %s:%v\n", callerName, r)
	}
}

func panicer(callerName string) {
	panic("panicer called from " + callerName)
}

func foo(wg *sync.WaitGroup) {
	defer wg.Done() //catch it here
	panicer("foo")
}

func PanicWithGoroutine() {
	defer handlePanic("main")
	var wg sync.WaitGroup
	wg.Add(1)
	//panic from one thread can be recovered only in that thread
	go func(wg *sync.WaitGroup) {
		// or catch it here
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println("recovered in func,", r)
			}
			// panic("panic from func")//not handled 
		}()
		foo(wg)
		fmt.Println("return from foo")
	}(&wg)
	wg.Wait()
	//panic("panic from main")//recovers in handle panic,below line not executed
	fmt.Println("Hello world")
}
