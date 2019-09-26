package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"brudnak_andrew":  []string{"Computer Science", "Reading to my Children", "Sleeping"},
		"brudnak_russell": []string{"Painting", "Reading", "Dancing"},
		"brudnak_ruby":    []string{"Toys", "Songs", "Eating Peas"},
	}

	m["brudnak_jessica"] = []string{"Netflix", "Making Videos", "Sleeping"}

	for k, v := range m {
		fmt.Printf("%v favorite things:\n", k)
		for i, val := range v {
			fmt.Printf("\t\tIndex: %v\t\t Value: %v\n", i, val)
		}
		fmt.Println()
	}

	fmt.Println(m)
}
