package main

import (
	"fmt"
)

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{"Andrew", 30}
	p2 := person{"Jessica", 30}
	p3 := person{"Russell", 2}
	p4 := person{"Ruby", 1}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
}
