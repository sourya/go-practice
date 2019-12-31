package main

import "fmt"

// typed constants
const hasKids bool = true
const yearBorn int = 1989

func main() {
	fmt.Println(hasKids)
	fmt.Println(yearBorn)
	fmt.Printf("Has kids: %T\n", hasKids)
	fmt.Printf("Year born: %T\n", yearBorn)
}
