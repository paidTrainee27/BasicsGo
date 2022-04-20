package main

import "fmt"

type Creature struct {
	species string
}

func (c Creature) reset() Creature {
	c.species = ""
	return c
}

func (c *Creature) resetWithPointer() {
	c.species = ""
}

func main() {
	onMethod()
}

func onMethod() {
	var creature *Creature = &Creature{species: "shark"}
	fmt.Printf("%+v\n", *creature)
	nCreate := creature.reset()
	creature.reset()
	fmt.Printf("%+v\n", nCreate)
	fmt.Printf("%+v\n", *creature)

}

func incrementValue(x *int) {
	*x++
}

func check1() {
	x := 0

	for {
		incrementValue(&x)
		if x >= 3 {
			break
		}
	}

	fmt.Println(x)
}
