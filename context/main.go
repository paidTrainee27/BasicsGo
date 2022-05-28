package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/sayhello", hello)
	fmt.Println("Started listening at 8089")
	http.ListenAndServe(":8089", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("hello request")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello World")
	// When you press ctrl + C
	case <-ctx.Done():

		fmt.Println("Req cancelled")
		err := ctx.Err()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
