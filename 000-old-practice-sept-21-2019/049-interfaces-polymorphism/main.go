package main

import "fmt"

type person struct {
	first string
	last  string
}

type child struct {
	person
	siblings bool
}

func (c child) speak() {
	fmt.Println("I am", c.first, c.last, " - the child speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case child:
		fmt.Println("I was pased into bar", h.(child).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	c1 := child{
		person: person{
			first: "Ruby",
			last:  "Brudnak",
		},
		siblings: true,
	}

	c2 := child{
		person: person{
			first: "Russell",
			last:  "Brudnak",
		},
		siblings: true,
	}

	p1 := person{
		first: "Andrew",
		last:  "Brudnak",
	}
	fmt.Println(c1)
	c1.speak()
	c2.speak()
	p1.speak()

	fmt.Println(p1)
	bar(c1)
	bar(c2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
