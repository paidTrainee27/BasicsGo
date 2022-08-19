package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	if err := assertStringType(123); err != nil {
		fmt.Println(err)
	}

}

func convInts() {
	var (
		index int8 = 15

		bigIndex int32
	)

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
		err = errors.New("type assertion failed")
	}
	return
}

func checkTypeCaste(typeCheck interface{}) {
	switch typeCheck.(type) {
	case bool:
		fmt.Println("type is boolean")
	case int:
		fmt.Println("type is integer")
	case string:
		fmt.Println("type is string")
	}
}

func all() {
	data := []byte(`{"k1":"val1","k2":"4","k3": ["apples","oranges",34]}`)
	var dataMap map[string]interface{}

	if err := json.Unmarshal(data, &dataMap); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", dataMap["k2"]) //string

	val := dataMap["k2"].(string) // from interface to string

	fmt.Println(reflect.TypeOf(dataMap["k2"])) //string

	val1, _ := strconv.ParseInt(val, 0, 8)

	fmt.Printf("%T %d\n", val1, val1) //returns int64 4

	val2, _ := strconv.Atoi(val)

	fmt.Printf("%T", val2) //int
}

