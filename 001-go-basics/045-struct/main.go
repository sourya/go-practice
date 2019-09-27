package main

import "fmt"

// COMPOSITE DATA STRUCTURE
// AGGREGATE DATA TYPES
// COMPOSE TOGETHER VALUES OF DIFFERENT TYPE
type person struct {
	first string
	last  string
	age   int
}

func main() {
	// CREATING a VALUE of TYPE person
	p1 := person{
		first: "Andrew",
		last:  "Brudnak",
		age:   29,
	}

	// CREATING a VALUE of TYPE person
	p2 := person{
		first: "Ruby",
		last:  "Brudnak",
		age:   1,
	}

	fmt.Println(p1.first)
	fmt.Println(p2.first)
}
