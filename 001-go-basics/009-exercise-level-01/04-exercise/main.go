package main

import "fmt"

var myInt int

type years int

var age years
func main() {
	age = 30
	fmt.Println(age)
	fmt.Printf("%T\n", age)

	// conversion not casting
	myInt = int(age)
	fmt.Println(myInt)
	fmt.Printf("%T\n", myInt)
}
