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

type adult struct {
	person
	hasKids bool
}

func main() {
	// CREATING a VALUE of TYPE person
	a1 := adult{
		person: person{
			first: "Andrew",
			last:  "Brudnak",
			age:   29,
		},
		hasKids: true,
	}

	// CREATING a VALUE of TYPE person
	p2 := person{
		first: "Ruby",
		last:  "Brudnak",
		age:   1,
	}

	fmt.Printf("%v is %v years old \t\t has kids: %v\n", a1.first, a1.age, a1.hasKids)
	fmt.Println(p2.first)
}
