package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main1() {
	goChan()
}

func readChannel(ch <-chan int) {
	// ch is read-only
}

func writeChannel(ch chan<- int) {
	// ch is write-only
}

func goChan() {
	orderChan := make(chan string)
	paymentChan := make(chan string)

	go shipOrders(orderChan)
	go makePayments(paymentChan)
	var exitOrder, exitPay bool

	for {
		//select/prints whichever chan has data without blobking the other chan
		select {
		case msg, open := <-orderChan:
			if !open {
				exitOrder = true
			} else {
				log(msg)
			}

		case msg, open := <-paymentChan:
			if !open {
				exitPay = true
			} else {
				log(msg)
			}
		}

		if exitOrder && exitPay {
			break
		}
	}
	//auto close
	// for msg := range orderChan {
	// 	log(msg)
	// }

	//Alternative
	// for {
	// 	msg, open := <-orderChan
	// 	if !open {
	// 		break
	// 	}
	// 	log(msg)
	// }

}

func shipOrders(out chan string) {
	defer close(out)

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second / 2)
		out <- fmt.Sprintf("prosessing %d", i)
	}
}

func makePayments(out chan string) {
	defer close(out)

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		out <- fmt.Sprintf("payment done %d", i)
	}
}

func goWaitRoutine() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		processOrders()
		wg.Done()
	}()
	wg.Wait()
}

func processOrders() {

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second / 2)
		log("processing order")
	}
}

func contextChan() {
	ctx := context.Background()
	cctx, cancel := context.WithCancel(ctx)
	cvtx := context.WithValue(cctx, "requestID", 1234)

	for val := range gen(cvtx) {
		fmt.Println(val)
		if val == 10 {
			break
		}
	}

	cancel()
	time.Sleep(1 * time.Second)

}

func gen(ctx context.Context) <-chan int {
	out := make(chan int)
	i := 0

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				fmt.Println("context value", ctx.Value("requestID"))
				close(out)
				return
			case out <- i: //?
				i++
			}
		}
	}()

	return out
}

func log(msg string) {
	fmt.Println(msg)
}
