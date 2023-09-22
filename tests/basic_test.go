package main

import "testing"

func TestSquare(t *testing.T) {
	n := 4
	e := 16
	if r := square(n); r != e {
		t.Errorf("Error testing func Square expected %d got %d", e, r)
	}
}

func TestSquareSlice(t *testing.T) {
	tc := []struct {
		name  string
		input int
		exp   int
		error error
	}{
		{"2 square", 2, 4, nil}, {"0 square", 0, 0, nil}, {"Negative square", -2, 0, nil},
	}
	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {
			if r := square(v.input); r != v.exp {
				t.Errorf("Error testing func Square expected %d got %d", v.exp, r)
			}
		})
	}

}
