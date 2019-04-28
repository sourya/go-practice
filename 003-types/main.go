package main

import "fmt"

var y = 62

//  DECLARED "z" as TYPE string
// interpreted string literal
var z = "Falcon 9’s first stage incorporates nine Merlin engines"

// raw string literal
var e = `Elon Musk said, "I think it is possible for ordinary
people to choose to be extraordinary."`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	// cannot reassign variable z to an int
	// since it has already been declared as string
	// z = 90
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)
}
