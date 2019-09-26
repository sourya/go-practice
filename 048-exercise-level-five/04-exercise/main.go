package main

import "fmt"

func main() {
	s := struct {
		first    string
		last     string
		age      int
		hasKids  bool
		hasCar   bool
		children []string
		vehicles map[string]string
	}{
		first:    "Andrew",
		last:     "Brudnak",
		age:      29,
		hasKids:  true,
		hasCar:   true,
		children: []string{"Russell", "Ruby"},
		vehicles: map[string]string{
			"SUV":   "Toyota",
			"Sedan": "Volkswagen",
		},
	}

	fmt.Printf("%v %v is %v years old. \tKIDS: %v \tCAR: %v\n", s.first, s.last, s.age, s.hasKids, s.hasCar)
	for i, v := range s.children {
		fmt.Printf("\t\tChild: %v \t Name: %v\n", i+1, v)
	}
	fmt.Println()
	for k, v := range s.vehicles {
		fmt.Printf("\t\tVehicle: %v \t Make: %v\n", k, v)
	}

}
