package main

import (
	"fmt"
	"sort"
)

//map[key-type]val-type
var dict1 map[string]int

func main() {

	// compute1()
	// sortMap()
	checkKey()
}

func sortWithoutSlice() {
	m := map[string]int{
		"bikram": 2,
		"akash":  10,
		"vijay":  5,
		"wasim":  2,
		"sahil":  2,
		"umesh":  2,
	}
	//to store previous state
	a := 0
	b := ""
l1: //can only be used above and inside loops
	for v, e := range m {
		c := v       //store current key ( if want to sort by value then store value)
		if b == "" { // Initialize
			a = e
			b = v
		} else {
			if b > c { //compare
				delete(m, b) // delete
				m[b] = a     //put at the end of map
				b = ""       //reset
				break l1     //break and restart
			}
		}
	}
	fmt.Println(m)
}

func declare() {
	//naive way 1
	dict := make(map[string]int)
	dict["first"] = 0
	dict["second"] = 1
	dict["third"] = 3
	delete(dict, "second")

	//way 2
	dict1 = map[string]int{"first": 0}

	fmt.Printf("%v", dict)
	fmt.Printf("%v", dict1)
}

func compute1() {

	dict := make(map[string]int)
	dict["first"] = 0
	dict["second"] = 1
	dict["third"] = 3

	el := dict["first"]

	if el > 10 {
		fmt.Println("> 10")
	} else {
		fmt.Println("else")
	}

	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	items := make([]string, len(sammy))

	var i int

	for _, v := range sammy {
		items[i] = v
		i++
	}
	fmt.Printf("%q", items)
}

func sortMap() {
	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, idx := range names {
		fmt.Println(idx, m[idx])
	}

}

func checkKey() {
	m1 := map[string]string{
		"name": "riyaz",
		"id":   "21",
	}
	if val, ok := m1["cut"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("not found")
	}
}

func deleting() {
	permissions := map[int]string{1: "read", 2: "write", 4: "delete", 8: "create", 16: "modify"}
	delete(permissions, 16)
	fmt.Println(permissions)
	//or
	permissions = map[int]string{}
}
