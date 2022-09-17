package main

import "fmt"

type cls1 func() int

var shadowCount int

var isEven func(num int) bool

/*
A function definition that is not bound to an identifier.
Anonymous functions are often arguments being passed to higher-order functions or
used for constructing the result of a higher-order function that needs to return a function.
If the function is only used once, or a limited number of times, an anonymous function may be
syntactically lighter than using a named function.
*/
func AnanymousFunc() {
	name := "Albert"

	// self Calling
	func(name string) {
		fmt.Println("Welcome" + name)
	}(name)

	//as a variable for minimal implementation
	isEven = func(num int) bool {
		return num%2 == 0
	}

	intArr := []int{1, 2, 3, 4}
	for i := 0; i < len(intArr); i++ {
		if isEven(intArr[i]) {
			fmt.Println(intArr[i])
		}
	}

}

func callClosure() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

func intSeq() cls1 {
	i := 0
	return func() int {
		i++ //captures the state of functions variable i
		return i
	}
}

func ShadowingVar() {
	//shadowing
	shadowCount := 2
	fmt.Println(shadowCount)
}

func closureVSanonymous() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i) //closures the state of variable i
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i) //anonymous func prints i as 3
			}()
		}
	}()
}
