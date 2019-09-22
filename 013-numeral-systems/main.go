package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]

	fmt.Println(n)
	fmt.Printf("decimal: \t\t%T\n", n)
	fmt.Printf("binary: \t\t%b\n", n)
	fmt.Printf("hexidecimal: \t%#X\n", n)
}
