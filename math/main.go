package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
}

func floatRound() {
	s := fmt.Sprintf("%.2f", 12.3456)
	x := 12.3456
	fmt.Println(math.Floor(x*100) / 100) // 12.34 (round down)
	fmt.Println(math.Round(x*100) / 100) // 12.35 (round to nearest)
	fmt.Println(math.Ceil(x*100) / 100)  // 12.35 (round up)
	fmt.Println(s)                       // s == "12.35"

}

func mod() {
	q := 36.0
	r := 8.0

	s := math.Mod(q, r)

	fmt.Println(s)
}

func randFunc() {
	// rand.Seed(time.Now().UnixNano()) //add at init
	target := rand.Intn(100)//number between 0-99
	fmt.Println(target)         
}
