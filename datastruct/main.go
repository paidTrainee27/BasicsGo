package main

import (
	"fmt"
	"sort"
)

var array1 [5]int

var users []user

type mSlice []string

type user struct {
	id   int
	name string
}

type userMap map[int]user

// arr := [...]int{2, 3, 8, 4} // A Array not a slice

func main() {
	// appendStructs()
	// mergeSortedArrays()
	// sortMap()
	twoDimensionArray()
}

func arrayDef() {
	var (
		arr  []int               //slice
		arr1 [5]int              //array
		arr2 = make([]int, 1, 2) //slice cannot create array
		arr3 = []int{}
		arr4 = [4]int{1, 2, 3, 5}
		arr5 = [...]int{1, 2} //is a array
	)
	arr2 = append(arr2, 4)
	fmt.Println(arr, arr1, arr2, arr3, arr4)
	arr2 = append(arr2, 1)
	arr2 = append(arr2, 1, 1, 1, 1)
	//cap  always  doubles the specified count(third arg) every time length exceeds the capacity.
	//len is size(first arg in make) + no. of element(append)
	fmt.Println("Hello World", len(arr2), cap(arr2))
	fmt.Println(arr5)
}

func appendStructs() {
	//logs.printLine(collectionBasic)
	array1[4] = 4

	slice1 := []int{1, 3, 5, 7}
	slice1[3] = 6
	// returns new COPY of slice.
	slice1 = append(slice1, 8)

	// fmt.Println(array1)
	fmt.Println(slice1)

	// users = append(users, userMap{1: user{1, "ryaaz"}})
	// users = append(users, userMap{2: user{2, "faraz"}})
	// users = append(users, userMap{3: user{3, "shayjaz"}})
	// users = append(users, userMap{4: user{4, "arbaaz"}})

	for _, user := range users {
		fmt.Printf("students %v\n", user)
	}
}

func mergeSortedArrays() {
	nums1 := []int{2, 4, 6, 8, 10}
	// m := 0
	nums2 := []int{1, 3, 5, 7, 9}
	// n := 5

	merge(&nums1, nums2)

	fmt.Println(nums1)

}

func merge(arr1 *[]int, arr2 []int) {
	for i := 0; i < len(arr2); i++ {
		if i%2 == 0 {
			continue
		}
		*arr1 = append(*arr1, arr2[i])
	}
	//OR
	//*arr1 = append(*arr1, arr2...)
	sort.Ints(*arr1)
}

func popIndex() {
	usersLice := mSlice{"A", "B", "C", "D", "asd", "asq"}
	usersLice.pop(0)
	fmt.Println(usersLice)
}

func (s *mSlice) pop(indx int) {
	//
	//*s[0] is interpreted by the compiler as *(s[0]). So first it gets the first element of the slice,
	//then applies the pointer deref. What we want is for it to first deref the slice, then get the element.
	//This can be done with (*s)[0]. It first applies the * to the slice and then indexes
	//
	(*s) = append((*s)[:indx], (*s)[indx+1:]...)
	// fmt.Println(s)
}

func sortMap() {
	users = append(users, user{5, "akroz"})
	users = append(users, user{3, "akram"})
	users = append(users, user{1, "armaan"})
	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].id < users[j].id
	})

	fmt.Println(fmt.Sprintf("%+v", users))
}

// Adding a element at the end of array
func push() []int {
	a := []int{1, 2, 3, 4, 5}
	v := 6
	return append(a, v)
}

// Removing a element from the end of array and returning that elem and slice
func pop() (int, []int) {
	a := []int{1, 2, 3, 4, 5}
	return a[len(a)-1], a[:len(a)-1] //the element and atray
}

// Removing a element from the start of array
func shift() (int, []int) {
	a := []int{1, 2, 3, 4, 5}
	return a[0], a[1:]
}

// adding a element at the start of array
func unshift() []int {
	a := []int{1, 2, 3, 4, 5}
	v := 0
	return append([]int{v}, a...)
}

func filterInPlace() []int {
	a := []int{1, 2, 3, 4, 5}

	//filtering logic/condition
	filterFunc := func(x int) bool {
		return x > 2
	}
	n := 0
	for _, x := range a {
		if filterFunc(x) {
			a[n] = x // or use new array
			n++
		}
	}
	return a[:n]
}

func appendSlice() []int {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	return append(a, b...)
}

func copySliceV1() []int {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func twoDimensionArray() {
	var twoDArray [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoDArray[j][i] = i + j
		}
	}
	fmt.Println(twoDArray)
}

type ServicePay struct {
	Name    string
	Payment int
}

var ServicePayList []ServicePay

func SortSliceOfStruct() {
	sp1 := ServicePay{
		Name:    "amaey",
		Payment: 43,
	}
	sp2 := ServicePay{
		Name:    "atul",
		Payment: 114,
	}
	sp3 := ServicePay{
		Name:    "ravi",
		Payment: 12,
	}
	ServicePayList = []ServicePay{sp1, sp2, sp3}
	sort.Slice(ServicePayList, comparePrice)
}
func comparePrice(i, j int) bool {
	return ServicePayList[i].Payment < ServicePayList[j].Payment
}
