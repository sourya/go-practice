package main

import "fmt"

func main() {
	ab := []string{"Andrew", "Brudnak", "Popcorn", "Green Tea"}
	fmt.Println(ab)
	rb := []string{"Russell", "Brudnak", "Humus", "Coconut Water"}
	fmt.Println(rb)

	// Slice of people OF slice of strings
	// Multi-Dimensional Slice
	xp := [][]string{ab, rb}
	fmt.Println(xp)
}
