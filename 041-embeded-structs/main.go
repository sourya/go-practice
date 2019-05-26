package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type child struct {
	person
	walking bool
}

func main() {
	p1 := person{
		first: "Andrew",
		last:  "Brudnak",
		age:   29,
	}

	c1 := child{
		person: person{
			first: "Ruby",
			last:  "Brudnak",
			age:   1,
		},
		walking: false,
	}

	fmt.Printf("First name: %v \t Last name: %v \t Age: %v \n", p1.first, p1.last, p1.age)
	fmt.Printf("First name: %v \t Last name: %v \t Age: %v \t Walking: %v \n", c1.first, c1.last, c1.age, c1.walking)
}
