package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type msgID string

const (
	requestId msgID = "msgId"
)

func main() {
	hlo := http.HandlerFunc(hello)
	http.Handle("/sayhello", hellowMiddleware(hlo))
	fmt.Println("Started listening at 8089")
	http.ListenAndServe(":8089", nil)
}

func hellowMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgId := time.Now().Nanosecond()
		ctx := context.WithValue(r.Context(), requestId, msgId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("hello request")
	fmt.Printf("msgId: %d\n", ctx.Value(requestId))

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
