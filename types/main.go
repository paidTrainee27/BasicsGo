package main

import (
	"errors"
	"fmt"
)

func main() {

	if err := assertStringType(123); err != nil {
		fmt.Println(err)
	}

}

func convInts() {
	var index int8 = 15

	var bigIndex int32

	bigIndex = int32(index)

	fmt.Printf("bigIndex data type: %T\n", bigIndex)

	var f float64 = 390.8
	var i int = int(f)

	fmt.Printf("f = %.2f\n", f)
	fmt.Printf("i = %d\n", i)

	a := 5 / 2
	fmt.Println(a) //prints 2
}

func convFloats() {
	var x int64 = 57

	var y float64 = float64(x)

	//round up to 2 decimal places
	fmt.Printf("%.2f\n", y)

	a := 5.0 / 2
	fmt.Println(a) //prints 2.5
}

func assertStringType(data interface{}) (err error) {
	if _, ok := data.(string); !ok {
		err = errors.New("Type assertion failed")
	}
	return
}
