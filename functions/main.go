package main

import "fmt"

type user struct {
	id   int
	name string
}

func main() {
	// callClosure()
	fmt.Println(getFactorial(4))
	fmt.Println(getFib(7))
}

func callVariacFunc() {
	u := user{
		id:   1,
		name: "Riaz",
	}
	callMee(1, 2, 3, 4, u)

}

//You can only have one variadic parameter in a function, and
//it must be the LAST parameter defined in the function.

//Variadic Generic function
func callMee(t ...interface{}) {
	for _, v := range t {
		fmt.Println(v)
	}
	// fmt.Println(t)
}

//Variadic function 
func callInt(i ...int) {
	fmt.Println(i)
}

func checkOrder() {
	names := []string{"Sammy", "Jessica", "Drew", "Jamie"}

	_ = join(",", "Sammy", "Jessica", "Drew", "Jamie")

	_ = join(",", "Sammy", "Jessica")
	_ = join(",", names...)
}

//Note changing the order of arg will give error : variadic args should be in the last
func join(del string, values ...string) string {
	var line string
	for i, v := range values {
		line = line + v
		if i != len(values)-1 {
			line = line + del
		}
	}
	return line
}
