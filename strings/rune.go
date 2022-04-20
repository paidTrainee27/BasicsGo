package main

import (
	"fmt"
	"unicode/utf8"
)

var myString string

func runMain() {
}

func checkString() {
	myString = "asda"

	for i := 0; i < len(myString); i++ {
		fmt.Println(myString[i])
	}

}

func checkRune() {
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
}
