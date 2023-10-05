package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// TASK :- if value of i equals 30 the exit with error
func main() {
	//initialization
	errGroup, ctx := errgroup.WithContext(context.Background())
	numberCh := make(chan int, 50)

	//goroutine 1
	errGroup.Go(func() error {
		defer close(numberCh)
		for i := 1; i <= 50; i++ {
			if i == 30 {
				return fmt.Errorf("30 is not allowed") // exiting: return error
			}
			numberCh <- i
			fmt.Printf("sending %d\n", i)
		}

		return nil
	})

	//catching error in another goroutine
	errGroup.Go(func() error {

		for j := 1; j <= 50; j++ {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case number := <-numberCh:
				fmt.Printf("receiving %d", number)
			}
		}
		return nil
	})

	//received error
	err := errGroup.Wait()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("done")
}
