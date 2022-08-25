package main

import (
	"fmt"
	"sync"
)

/*
It is important to pass the pointer of wg in line no. 21. If the pointer is not passed, then each Goroutine
will have its own copy of the WaitGroup and main will not be notified when they finish executing.
*/

func callLock() {
	i := 0
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(2)
	go func(c *int, w *sync.WaitGroup, mut *sync.Mutex) {
		go incOdd(c, w, mut)
		go incEven(c, w, mut)
	}(&i, &wg, &m)
	wg.Wait()
	fmt.Println(i)
}

func incOdd(inx *int, w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*inx++
	m.Unlock()
	w.Done()
}

func incEven(inx *int, w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*inx++
	m.Unlock()
	w.Done()
}

var x1 = 0

// Without mutex the value of x1 will be random due to race condition
func inxrement(wg *sync.WaitGroup, m *sync.Mutex) {

	m.Lock()
	x1 = x1 + 1
	m.Unlock()
	wg.Done()
}
func DeadLockMutext() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go inxrement(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
