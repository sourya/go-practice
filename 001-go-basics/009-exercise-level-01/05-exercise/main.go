package main

import "fmt"

var myInt int

type years int

var age years

var ageInt int
var ageYears years

func main() {
	age = 30
	fmt.Println(age)
	fmt.Printf("%T\n", age)

	// conversion not casting
	myInt = int(age)
	fmt.Println(myInt)
	fmt.Printf("%T\n", myInt)

	ageInt = 2
	fmt.Println(ageInt)
	fmt.Printf("%T\n", ageInt)

	fmt.Println(ageYears)
	fmt.Printf("%T\n", ageYears)

	ageYears = years(ageInt)
	fmt.Println(ageYears)
	fmt.Printf("%T\n", ageYears)
}
