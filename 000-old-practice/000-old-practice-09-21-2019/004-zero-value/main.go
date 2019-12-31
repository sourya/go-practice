package main

import "fmt"

var vehModel string
var modelYear int

func main() {
	fmt.Println("Printing the value of vehModel:", vehModel)
	fmt.Printf("%T\n", vehModel)

	vehModel = "Model Y"

	fmt.Println("Printing the value of vehModel:", vehModel)
	fmt.Printf("%T\n", vehModel)

	fmt.Printf("%T\n", modelYear)
}
