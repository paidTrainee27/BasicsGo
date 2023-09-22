package main

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

/*
	A Go string is a read-only slice of bytes.
	The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
	In other languages, strings are made of “characters”.
	In Go, the concept of a character is called a rune -
	it’s an integer that represents a Unicode code point
	Go string literals are UTF-8 encoded text.
*/

// var salutation string = "Hi there !"
/*WON'T WORK HERE*/
// salutation  := "Hi there !"

//stirng literals are of two types raw and interpreted
var backTick = `Say "hello" to Go!
It's a fast, statically typed,
compiled language that feels like a dynamically typed,
interpreted language.`

var stringLitrals = "Say \"hello\" to Go!" //interpreted

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

func urlParsing() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
