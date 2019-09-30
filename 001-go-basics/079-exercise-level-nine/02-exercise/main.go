package main

import "fmt"

type person struct {
	first string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("My name is %v, age %v\n", p.first, p.age)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Andrew",
		age:   29,
	}

	p2 := person{
		first: "Russell",
		age:   2,
	}

	saySomething(&p1)
	saySomething(&p2)

	p1.speak()
	p2.speak()

	// *********************
	// Code below won't run:
	// saySomething(p1)
	// *********************
}
