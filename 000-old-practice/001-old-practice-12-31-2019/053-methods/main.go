package main

import "fmt"

type person struct {
	first string
	last  string
}

type adult struct {
	person
	hasKids bool
}

func (a adult) speak() {
	fmt.Println("I am", a.first, a.last)
}

func main() {
	a1 := adult{
		person: person{
			first: "Andrew",
			last:  "Brudnak",
		},
		hasKids: true,
	}

	a2 := adult{
		person: person{
			first: "Jessica",
			last:  "Brudnak",
		},
		hasKids: true,
	}

	fmt.Println(a1)

	a1.speak()
	a2.speak()
}
