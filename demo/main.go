package main

import (
	"fmt"
	"piscine"
)

func main() {
	quadSelect := "A"
	x := 0
	y := 0
	fmt.Printf("Select your type of Quad: A/B/C/D/E ")
	fmt.Scanln(&quadSelect)
	if quadSelect < "A" || quadSelect > "E" {
		fmt.Println("Please select an existing type of Quad.")
		return
	}
	fmt.Printf("x: ")
	fmt.Scanln(&x)
	fmt.Printf("y: ")
	fmt.Scanln(&y)
	fmt.Println()
	fmt.Printf("Here is a type %s Quad with a width of: %v and a height of %v: \n\n", string(quadSelect), x, y)
	if quadSelect == "A" {
		piscine.QuadA(x, y)
	} else if quadSelect == "B" {
		piscine.QuadB(x, y)
	} else if quadSelect == "C" {
		piscine.QuadC(x, y)
	} else if quadSelect == "D" {
		piscine.QuadD(x, y)
	} else if quadSelect == "E" {
		piscine.QuadE(x, y)
	}
}
