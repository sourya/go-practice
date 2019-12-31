package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Andrew",
		last:  "Brudnak",
		age:   29,
	}

	p2 := person{
		first: "Ruby",
		last:  "Brudnak",
		age:   1,
	}

	fmt.Printf("First name: %v \t Last name: %v \t Age: %v \n", p1.first, p1.last, p1.age)
	fmt.Printf("First name: %v \t Last name: %v \t Age: %v \n", p2.first, p2.last, p2.age)
}
