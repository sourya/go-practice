package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "Andrew",
		last:  "Brudnak",
		favFlavors: []string{
			"Chocolate", "Green Tea", "Vanilla",
		},
	}

	p2 := person{
		first: "Russell",
		last:  "Brudnak",
		favFlavors: []string{
			"Chocolate", "Birthday Cake", "Cherry Garcia",
		},
	}
	fmt.Printf("%v\"s favorite ice cream flavors:\n", p1.first)
	for _, v := range p1.favFlavors {

		fmt.Printf("\t\t%v\n", v)
	}

	fmt.Println()

	fmt.Printf("%v\"s favorite ice cream flavors:\n", p2.first)
	for _, v := range p2.favFlavors {

		fmt.Printf("\t\t%v\n", v)
	}
}
