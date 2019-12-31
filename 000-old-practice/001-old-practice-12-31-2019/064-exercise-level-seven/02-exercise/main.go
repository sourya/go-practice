package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "Ruby",
	}
	fmt.Println(p1)
	changeMe(&p1, "Russell")
	fmt.Println(p1)
}

func changeMe(p *person, s string) { p.name = s }
