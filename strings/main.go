package main

import (
	"fmt"
	"strconv"
	"strings"
)

// var salutation string = "Hi there !"

/*WON'T WORK HERE*/
// salutation  := "Hi there !"
var backTick = `Say "hello" to Go!
It's a fast, statically typed,
compiled language that feels like a dynamically typed,
interpreted language.`

var stringLitrals = "Say \"hello\" to Go!"

func main() {
	salutation := "Hi there !"
	fmt.Println(salutation)

}

func metthods() {
	ss := "Sammy Shark"
	tt := []string{"sharks", "crustaceans", "plankton"}
	fmt.Println(strings.ToUpper(ss))
	fmt.Println(strings.ToLower(ss))
	fmt.Println(strings.HasPrefix(ss, "Sammy"))
	fmt.Println(strings.HasSuffix(ss, "Shark"))
	fmt.Println(strings.Contains(ss, "Sh"))
	fmt.Println(strings.Count(ss, "S"))
	fmt.Println(len(ss))
	fmt.Println(strings.Join(tt, ","))
	fmt.Println(strings.Split(ss, " "))

	data := "  username password     email  date"
	fields := strings.Fields(data)
	fmt.Printf("%q", fields) //add quotes to each word

	fmt.Println(strings.ReplaceAll(ss, "Shark", "Leopard"))
	ss = strings.TrimSpace(ss)

}

func splitStringToArray(s string) {
	params := strings.Split(s, "/")
	productId, err := strconv.Atoi(params[len(params)-1])
	//Also ..but returns int64
	if i, err := strconv.ParseInt(params[len(params)-1], 10, 64); err == nil {
		fmt.Printf("i=%d, type: %T\n", i, i)
	}

	fmt.Println(productId, err)
}

func intToString() {
	i := 0
	str := fmt.Sprintf("i is:%d", i)
	//or
	str = fmt.Sprint(i)

	fmt.Print(str)

	a := strconv.Itoa(12)
	fmt.Printf("%q\n", a)

}

func floatToString() {
	f := 5524.53
	fmt.Println("Sammy has " + fmt.Sprint(f) + " points.")
}

func fromString() {
	s := "yourDomain/path/1"
	params := strings.Split(s, "/")
	productId, err := strconv.Atoi(params[len(params)-1])
	//Also ..but returns int64
	if i, err := strconv.ParseInt(params[len(params)-1], 10, 64); err == nil {
		fmt.Printf("i=%d, type: %T\n", i, i)
	}

	fmt.Println(productId, err)

	/**IMP**/
	a := "not a number"
	b, err := strconv.Atoi(a)
	fmt.Println(b)   //prints 0
	fmt.Println(err) //error

	a = "my string"
	e := []byte(a)
	fmt.Println(e)
}

func toBytes() {
	a := "my string"

	b := []byte(a)

	c := string(b)

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)
}
