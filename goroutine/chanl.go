package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// withWaitGroup()
	// withChannel()
	// multiReceiverChannel()
	blockingChannel()
}


func blockingChannel()  {
	 ch := make(chan bool)

	 go func (chan bool)  {
		 defer close(ch)
		 fmt.Print("waiting...")
		 ch <- true
		 fmt.Println("Done")
	 }(ch)

	 <-ch
	 fmt.Println("exiting.")

}

func multiPrint(wg *sync.WaitGroup, c <-chan int, rx int) {
	defer func() {
		fmt.Println("exiting Rx", rx)
		wg.Done()
	}()
	for d := range c {
		// fmt.Println(fmt.Printf("Receiver %d received %d",rx,d))
		fmt.Printf("%d: read %d from channel\n", rx, d)
	}
}

func multiReceiverChannel() {
	ch := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go multiPrint(&wg, ch, i)
	}
	generateCh(ch)
	wg.Wait()
	fmt.Println("exiting main....")
}

func withChannel() {
	c1 := make(chan int)
	go printCh(c1)
	generateCh(c1)
	// time.Sleep(time.Second * 2)
}

//sender
func generateCh(c chan<- int) {
	defer func() {
		fmt.Println("closing channel")
		close(c)
	}()
	for i := 0; i < 3; i++ {
		c <- i
		time.Sleep(time.Second / 2)
	}
}

//receiver
func printCh(c chan int) {
	for i := range c {
		fmt.Println("data from channel ", i)
	}
}

func withWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)

	go generateNumbers(&wg, 3) //IMPORTANT to send pointer
	go printNumbers(&wg, 3)    //IMPORTANT to send pointer

	wg.Wait()
	fmt.Println("exiting main...")

}

func generateNumbers(wg *sync.WaitGroup, count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		time.Sleep(time.Second / 4)
		fmt.Println("generating numbers ", i)
	}
}

func printNumbers(wg *sync.WaitGroup, count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("printing numbers ", i)
	}

}
