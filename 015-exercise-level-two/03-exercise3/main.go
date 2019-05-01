package main

import "fmt"

const (
	age1 = 19 // untyped constant - of kind
	age2 int = 29 // typed constant - of type
)

func main() {
	fmt.Println(age1, age2)
}
