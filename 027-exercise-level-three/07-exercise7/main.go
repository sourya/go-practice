package main

import "fmt"

func main() {
	x := "Russell"

	if x == "Russell" {
		fmt.Println(x)
	} else if x == "Andrew" {
		fmt.Println("Excellent", x)
	} else {
		fmt.Println("Nope, neither")
	}
}
