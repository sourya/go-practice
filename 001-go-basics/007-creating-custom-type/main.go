package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 19
	b = 100
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
