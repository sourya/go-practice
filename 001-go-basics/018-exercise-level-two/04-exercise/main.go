package main

import "fmt"

func main() {
	num := 1
	fmt.Printf("decimal: %d\t binary: %b\t hex: %#x\n", num, num, num)
	// bit shifting
	num2 := num << 1
	fmt.Printf("decimal: %d\t binary: %b\t hex: %#x\n", num2, num2, num2)
}
