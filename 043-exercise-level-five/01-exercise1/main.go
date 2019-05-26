package main

import "fmt"

type person struct {
	first           string
	last            string
	favoriteFlavors []string
}

func main() {
	p1 := person{
		first: "Andrew",
		last:  "Brudnak",
		favoriteFlavors: []string{
			"Chocolate",
			"Black Cherry",
			"Strawberry",
		},
	}

	for i, v := range p1.favoriteFlavors {
		i++
		fmt.Printf("%v's favorite ice cream flavors: \t %v - %v\n", p1.first, i, v)
	}
}
