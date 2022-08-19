package main

import (
	"errors"
	"fmt"
	"math"
)

type MyError struct{}

type RequestError struct {
	StatusCode int

	Err error //error.New("")
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func (m *MyError) Error() string {
	return "Something went wrong"
}

func main() {
	customError()
}

func simpleError() {
	if resl, err := findSquareRoot(-5); err != nil {
		fmt.Println("There was an error:", err)
	} else {
		fmt.Printf("%.2f", resl)
	}

}

func findSquareRoot(num int) (result float64, err error) {
	if num < 0 {
		err = errors.New("value can't be less then zero")
		return
	}
	result = math.Sqrt(float64(num))
	return
}

func sayHello() (string, error) {
	return "", &MyError{}
}

func customError() {
	s, err := sayHello()
	if err != nil {
		fmt.Println("unexpected error: err:", err)
	}
	fmt.Println("The string:", s)
}

func doRequest() error {
	return &RequestError{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}

func customErrorStruct() {
	err := doRequest()
	if err != nil {
		fmt.Println(err)
	}
}
