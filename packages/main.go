package main

import (
	"fmt"
	_ "go-basics/packages/foo"
	_ "go-basics/packages/zoo"
)

var Global string = func() string {
	fmt.Println("main global initialisation")
	return "main global"
}()

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("Hello world")
}
