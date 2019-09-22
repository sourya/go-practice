package main

import "fmt"

type age int

var x age
var y int

func main() {
	fmt.Println("zero value of type age:", x)
	fmt.Printf("The type of variable 'x': %T\n", x)
	x = 19
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("The type of variable 'y': %T\n", y)

}
