package types

type student struct {
	marks int
}

func CheckRefOrValType() {
	i := 0
	std := student{}
	m := map[string]int{
		"count": 0,
	}
	arr1 := [...]int{1, 2, 3}
	sl := []int{1, 2, 3}

	check(&i, &std, &arr1, m, sl)
}

/*
	We need to pass pointer to update the values of basic data types like
	int, float, string, bool, struct, array are called value types.

	while for map slice pointer channel and function without pointer we can update the value
	are called reference types.
*/

func check(i *int, s *student, r *[3]int, m map[string]int, sl []int) {
	s.marks++ //struct
	r[0] = 10 //array
	*i++      // int

	m["count"]++ //map
	sl[0] = 11   //slice
}
