package main

import "fmt"

// Constants of a kind
const a int = 19
const b float64 = 19.5
const c string = "Roland Deschain"

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
