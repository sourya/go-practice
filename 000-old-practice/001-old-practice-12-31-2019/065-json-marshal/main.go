package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "Ruby",
		Last:  "Brudnak",
		Age:   1,
	}

	p2 := Person{
		First: "Russell",
		Last:  "Brudnak",
		Age:   2,
	}

	people := []Person{
		p1,
		p2,
	}

	fmt.Println(people)

	b, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
