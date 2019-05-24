package main

import "fmt"

func main() {
	ab := []string{"Andrew", "Brudnak", "Lambda", "Go"}
	fmt.Println(ab)
	rb := []string{"Russell", "Brudnak", "Trolls", "Painting"}
	fmt.Println(rb)

	// Like a spreadsheet
	xp := [][]string{ab, rb}
	fmt.Println(xp)
}
