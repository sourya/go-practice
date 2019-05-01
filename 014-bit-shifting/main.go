package main

import "fmt"

const (
	_ = iota // 0 place iota thrown away
	kb = 1 << (iota * 10) // 10
	mb = 1 << (iota * 10) // 20
	gb = 1 << (iota * 10) // 30
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
