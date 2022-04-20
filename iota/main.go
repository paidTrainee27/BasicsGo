package main

import "fmt"

const (
	C0 = iota
	C1
	C2
)
const (
	B1 = iota + 1
	B2
	B3
)

func main() {
	// printSimple()
	printDirection(North)
	printDirection(South)

}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d *Direction) String() string {
	// fmt.Println(d)
	return []string{"North", "East", "South", "West"}[*d]
}

func printSimple() {
	fmt.Println(C0, C1, C2)
}

func printDirection(d Direction) {
	fmt.Print(&d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}
