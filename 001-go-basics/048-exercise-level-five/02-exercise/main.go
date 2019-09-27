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

	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	for _, v := range m {
		fmt.Printf("%v %v\"s favorite ice cream flavors:\n", v.first, v.last)
		for i, val := range v.favFlavors {
			fmt.Printf("\t\t#%v: %v\n", i+1, val)
		}
	}

}
