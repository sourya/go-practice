package main

import "fmt"

func main() {
	name := "Andrew"
	fmt.Println("value stored on name:", name)
	fmt.Println("name variable address in memory:", &name)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", &name)

}
