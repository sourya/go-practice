package main

import "fmt"

func main() {
	x := map[string][]string{
		"brudnak_andrew": {
			"Computer Science",
			"Playing with Russell & Ruby",
			"Spending time with Jessie",
		},
		"brudnak_russell": {
			"Watching YouTube",
			"Painting",
			"Eating fruit",
		},
	}

	for i, v := range x {
		fmt.Printf("%v\n", i)
		for index, favorite := range v {
			index++
			fmt.Printf("\t List number: %v \t Favorite thing: %v \n ", index, favorite)
		}
	}
}
