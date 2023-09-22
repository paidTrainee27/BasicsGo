package main

import (
	"errors"
	"fmt"
	"go-basics/logs"
	"sort"
)

type student struct {
	id      int
	name    string
	address address
}

type address string

func (a address) validate() (err error) {
	if a == "" {
		err = errors.New("student address cant be empty")
	}
	return
}

//Stringer
func (s student) String() string {
	return fmt.Sprintf("%s with id %d and address %s", s.name, s.id, s.address)
}

func (s *student) updateAddress(address address) {
	s.address = address
}

func (s student) getAddress() address {
	return s.address
}

func main() {
	// createStruct()
	// anonymousStruct()
	updateStruct()
}

func updateStruct() {
	user1 := &student{}
	user1.id = 2
	user1.name = "aayan"
	user1.address = "street back"
	fmt.Println("student created as", *user1)
	user1.updateAddress("New street front")
	fmt.Println("student updates as", *user1)
}

func createStruct() {
	user1 := &student{}
	user1.id = 1
	user1.name = "ryaaz"
	user1.address = "street back"

	if err := user1.address.validate(); err != nil {
		logs.PrintError(fmt.Errorf("Error %s", err))
	} else {
		fmt.Println("student created as", *user1)
	}
}

func anonymousStruct() {
	user2 := []struct {
		id      int
		name    string
		address address
	}{{1, "Mr.X", " Wall street"}, {1, "Mr.X", " Wall street"}, {1, "Mr.X", " Wall street"}}

	logs.PrintJson(user2)
}

func SortSliceOfStruct() {
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family)
}

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func SortSliceOfStructCustom() {

	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}

func SortDecensing() {
	arr := [...]int{12, 331, 31, 451, 54, 52} // unsorted
	arrSlice := arr[:]                        // created a slice of the array
	sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))

	/* OR
	sort.Ints(arrSlice)
	for i, j := 0, len(arrSlice)-1; i < j; i, j = i+1, j-1 {
		arrSlice[i], arrSlice[j] = arrSlice[j], arrSlice[i]
	}
	*/
	fmt.Println("sorted array in descending order:", arrSlice)
}
