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
	fmt.Println("I am", a.first, a.last, "- the adult speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar FROM SWITCH", h.(person).first)
	case adult:
		fmt.Println("I was passed into bar FROM SWITCH", h.(adult).first)
	}
	fmt.Println("I was passed into bar", h)
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

	p1 := person{
		first: "Shelby",
		last:  "Diamond",
	}

	fmt.Println(a1)

	a1.speak()
	a2.speak()

	fmt.Println(p1)

	p1.speak()

	bar(a1)
	bar(a2)
	bar(p1)
}
