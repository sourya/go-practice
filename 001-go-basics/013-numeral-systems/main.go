package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)

	n := bs[0]
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#x\n", n)
}
