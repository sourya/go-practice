package main

import "fmt"

var num = 19

func main() {
	fmt.Println("Variable 'num':", num)
	fmt.Printf("Type: %T\n", num)
	fmt.Printf("Binary: %b\n", num)
	fmt.Printf("Hexidecimal: %x\n", num)
	fmt.Printf("Hex with zero x: %#x\n", num)
}
