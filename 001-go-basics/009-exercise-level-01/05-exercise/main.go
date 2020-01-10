package main

import "fmt"

var myInt int

type years int

var age years


var age2 int
var age3 years

func main() {
	age = 30
	fmt.Println(age)
	fmt.Printf("%T\n", age)

	// conversion not casting
	myInt = int(age)
	fmt.Println(myInt)
	fmt.Printf("%T\n", myInt)

	age3 := 10
	fmt.Println(age3)
	fmt.Printf("%T\n", age3)

	fmt.Println(age2)
	fmt.Printf("%T\n", age2)

	age2 := years(age3)
	fmt.Println(age2)
	fmt.Printf("%T\n", age2)
}
