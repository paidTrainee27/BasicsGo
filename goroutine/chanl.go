package goroutine

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

/*
 An unbuffered channel is used to perform synchronous communication between goroutines while a buffered channel is
 used for perform asynchronous communication. An unbuffered channel provides a guarantee that an exchange between
 two goroutines is performed at the instant the send and receive take place.
 A buffered channel has no such guarantee.
*/

/*
In buffered channel there is a capacity to hold one or more values before they're received. In this types of
channels don't force goroutines to be ready at the same instant to perform sends and receives. There are also
different conditions for when a send or receive does block. A receive will block only if there's no value in the
channel to receive. A send will block only if there's no available buffer to place the value being sent.
*/
func BufferedAsynchoronousCh() {
	ch1 := make(chan int, 1)
	ch3 := make(chan struct{})
	go genFunc(ch1, ch3)
	go revFunc(ch1, ch3)
	<-ch3
	fmt.Println("Hello World")

}

/*
In unbuffered channel there is no capacity to hold any value before it's received.
In this type of channels both a sending and receiving goroutine to be ready at the same instant before any send
or receive operation can complete. If the two goroutines aren't ready at the same instant, the channel makes the
goroutine that performs its respective send or receive operation first wait. Synchronization is fundamental in
the interaction between the send and receive on the channel. One can't happen without the other.
*/
func UnbufferedsynchoronousCh() {
	ch1 := make(chan int)
	ch3 := make(chan struct{})
	go genFunc(ch1, ch3)
	go revFunc(ch1, ch3)
	<-ch3
	fmt.Println("Hello World")

}

func genFunc(c chan int, c3 chan struct{}) {
	defer func() {
		fmt.Println("exiting genFunc")
		close(c3)
	}()
	fmt.Println("printing b4")
	c <- 1
	fmt.Println("printing genFunc")
}
func revFunc(c chan int, c3 chan struct{}) {
	// comment defer block for unbuffered channel to print "printing genFunc"
	defer func() {
		fmt.Println("exiting revFunc")
		close(c3)
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("after sleep")
	fmt.Println(<-c)
}

func QuitAChannel() {

	quit := make(chan bool)
	go func() {
		defer func() {
			fmt.Println("Returning")
		}()
		for {
			select {
			case <-quit:
				return
				// do other stuff
			}
		}
	}()
	// Do stuff
	// quit goroutine
	// time.Sleep(time.Second / 2)
	quit <- true
	fmt.Println("Finally")
}

func blockingChannel() {
	ch := make(chan bool)

	go func(chan bool) {
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

func MultiReceiverChannel() {
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

func WithChannel() {
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

func WithWaitGroup() {
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

func buffChan() {
	var wg sync.WaitGroup
	ch := make(chan int, 2)
	wg.Add(1)
	go func(c chan int, w *sync.WaitGroup) {
		i := 17
		c <- i
		c <- i + 18
		defer func() {
			close(c) //imp if using range
			w.Done()
		}()

	}(ch, &wg)
	wg.Wait()
	// close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func buffSimple() {
	//for writing and reading from same function buffer is imp or else error
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "world"

	fmt.Printf("Lenth of Channel: %d \n", len(ch))

	fmt.Printf("Prinf %s \n", <-ch)

	close(ch)

	for v := range ch {
		fmt.Printf("Range %s \n", v)
		// fmt.Println(v)
	}
}

func LimitMaxGoRoutines() {
	//Declaration
	var i int
	var wg sync.WaitGroup
	var m sync.Mutex

	//Initialization
	ch := make(chan struct{}, 10)

	for j := 1; j <= 10; {
		for k := 1; k <= 10; k++ {
			wg.Add(1) // note the place
			ch <- struct{}{}
			go incCount(&i, &wg, &m, ch)
			fmt.Println("inner loop", k)
		}
		fmt.Println("Outer loop", j)
		j++
	}

	wg.Wait()

	fmt.Println("final value", i)
}

func incCount(i *int, wg *sync.WaitGroup, m *sync.Mutex, c chan struct{}) {
	defer func() {
		wg.Done()
		<-c
	}()
	m.Lock()
	*i++
	m.Unlock()

}

/*
Since sending blocks if there is no reader and reading is blocking if there's no sender, and you're a. waiting for
both goroutines to finish first and b. doing one more read than writes (the Println read), you need a buffered
channel, that has exactly one extra place in the buffer.

You need to push an initial value in the channel for the process to start.
*/

func increment(c chan int, wg *sync.WaitGroup) {
	for x := 0; x < 10; x++ {
		a := <-c
		fmt.Println("increment read ", a)
		a++
		c <- a
	}
	fmt.Println("Incrment done!")
	wg.Done()
}

func decrement(c chan int, wg *sync.WaitGroup) {
	for x := 0; x < 10; x++ {
		a := <-c
		fmt.Println("Decrement read ", a)
		a--
		c <- a
	}
	fmt.Println("Dencrment done!")
	wg.Done()
}

var x = 0

func SharedResourcesWithChannel() {
	// GOMAXPROCS(NumCPU())

	//we create a buffered channel with 1 extra space. This means
	//you can send one extra value into it if there is no reader, allowing for the final result to be pushed to println
	c := make(chan int, 1)

	//we create a wait group so we can wait for both goroutines to finish before reading the result
	wg := sync.WaitGroup{}
	wg.Add(1) //mark one started
	go increment(c, &wg)
	wg.Add(1) //mark another one started. We can just do Add(2) BTW
	go decrement(c, &wg)

	//now we push the initial value to the channel, starting the dialog
	c <- x

	//let's wait for them to finish...
	wg.Wait()

	//now we have the result in the channel's buffer
	fmt.Println("Total: ", <-c)
}
