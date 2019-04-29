package main

import "fmt"

var y = 19

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // type
	fmt.Printf("%b\n", y) // binary
	fmt.Printf("%x\n", y) // hexadecimal
	fmt.Printf("%#x\n", y) // hexadecimal with zero x

	y = 1

	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y) // "\t" escape tab

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) // string version
	fmt.Println(s)
	fmt.Printf("%v", y) // normal value
}