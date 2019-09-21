package main

import "fmt"

var message string

// USE var for zero value and package level scope

func main() {
	fmt.Println("Printing the value of 'message' variable:", message)
	fmt.Printf("%T\n", message)
	message = "Welcome back anonymous user"
	fmt.Println("Printing the value of 'message' variable:", message)
	fmt.Printf("%T", message)
}
