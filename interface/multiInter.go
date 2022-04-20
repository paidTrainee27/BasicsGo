package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Square struct {
	Width  float64
	Height float64
}

func (c *Circle) Area() (area float64) {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) String() string  {
	return fmt.Sprintf("Area is %.2f",c.Area())
}

func (s *Square) Area() (area float64) {
	return s.Width * s.Height
}

type Sizer interface {
	Area() float64
}

//A way to implement multiple interfaces on one struct
type Shaper interface {
	Sizer
	fmt.Stringer
}

func RunMultipleInterfaces() {
	var c Shaper = &Circle{2.3}
	var s Sizer = &Square{3,2}
	fmt.Printf("Type is %T \n", s)
	fmt.Printf("Area is %.2f \n", s.Area())
	fmt.Println(c.String())

}

// func main()  {
// runMultipleInterfaces()
// }
