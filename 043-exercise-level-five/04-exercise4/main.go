package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		last      string
		age       int
		children  map[int]string
		languages []string
		married   bool
	}{
		first: "Andrew",
		last:  "Brudnak",
		age:   29,
		children: map[int]string{
			1: "Russell",
			2: "Ruby",
		},
		languages: []string{
			"JavaScript",
			"Go",
			"Python",
			"Rust",
		},
		married: true,
	}

	fmt.Printf("%v %v is %v years old and has %v childen.\n" +
		"Their names are %v and %v\n" +
		"His favorite pogramming language is %v", p1.first, p1.last, p1.age,
	len(p1.children), p1.children[1], p1.children[2], p1.languages[1])
}
