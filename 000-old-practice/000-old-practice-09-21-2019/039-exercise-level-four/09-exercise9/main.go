package main

import "fmt"

func main() {
	x := map[string][]string{
		"brudnak_andrew": []string{
			"Computer Science",
			"Playing with Russell & Ruby",
			"Spending time with Jessie",
		},
		"brudnak_russell": []string{
			"Watching YouTube",
			"Painting",
			"Eating fruit",
		},
	}

	x["brudnak_ruby"] = []string{
		"Being held by mommy",
		"Seeing new things",
		"Talking with my parents",
	}

	for i, v := range x {
		fmt.Printf("%v\n", i)
		for index, favorite := range v {
			index++
			fmt.Printf("\t List number: %v \t Favorite thing: %v \n ", index, favorite)
		}
	}
}
