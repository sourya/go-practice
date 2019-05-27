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

	people := map[string]person{
		p1.last: p1,
	}

	for _, v := range people {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favoriteFlavors {
			i++
			fmt.Println(i, val)
		}
		fmt.Println("----------")
	}
}
