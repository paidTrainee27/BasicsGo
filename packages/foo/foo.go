package foo

import "fmt"

var FooGlobal string = func() string {
	fmt.Println("foo global initialisation")
	return "foo global"
}()

func init() {
	fmt.Println("foo init")
}

func Bar() {
	fmt.Println("This function lives in an another file!")
}