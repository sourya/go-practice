package main

import "fmt"

var a int
var b float64
var cMax int8 = 127

func main() {
	x := 19
	y := 19.2
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	a = 1
	b = 1.2

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	b = 12
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	fmt.Println(cMax)
}
