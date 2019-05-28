package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v, I am %v years old.", p.first, p.age)
}

func main() {
	p1 := person{
		first: "Andrew",
		last:  "Brudnak",
		age:   29,
	}

	p1.speak()
}
