package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"brudnak_andrew":  {"Computer Science", "Reading to my Children", "Sleeping"},
		"brudnak_russell": {"Painting", "Reading", "Dancing"},
		"brudnak_ruby":    {"Toys", "Songs", "Eating Peas"},
	}

	delete(m, "brudnak_andrew")

	for k, v := range m {
		fmt.Printf("%v favorite things:\n", k)
		for i, val := range v {
			fmt.Printf("\t\tIndex: %v\t\t Value: %v\n", i, val)
		}
		fmt.Println()
	}

	fmt.Println(m)
}
