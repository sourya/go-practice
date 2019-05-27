package main

import "fmt"

type person struct {
	first string
	last  string
}

type child struct {
	person
	siblings bool
}

func (c child) speak() {
	fmt.Println("I am", c.first, c.last)
}

func main() {
	c1 := child{
		person: person{
			first: "Ruby",
			last:  "Brudnak",
		},
		siblings: true,
	}

	c2 := child{
		person: person{
			first: "Russell",
			last:  "Brudnak",
		},
		siblings: true,
	}
	fmt.Println(c1)
	c1.speak()
	c2.speak()
}
