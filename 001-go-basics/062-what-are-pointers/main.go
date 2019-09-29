package main

import "fmt"

func main() {
	a := 42
	fmt.Println("Value of variable 'a' is:", a)
	// AND what is the address
	// AND what is the address
	fmt.Println("Address of where the value is stored for 'a' is:", &a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// EVERYTHING in Go is pass by VALUE

	b := &a
	fmt.Println(&a)
	fmt.Println(b)

	fmt.Println(*b)
	fmt.Println(a)

	a = 19

	// & gives you the address
	// * gives you the value stored at the address

	fmt.Println(*b)
	fmt.Println(a)

	*b = 2020

	fmt.Println(*b)
	fmt.Println(a)

}
