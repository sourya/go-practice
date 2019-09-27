package main

import "fmt"

var x int
var y string
var z bool

// zero values being printed
func main() {
	fmt.Println("int\t\t\t\tzero value:", x)
	fmt.Println("string\t\t\tzero value:", y)
	fmt.Println("bool\t\t\tzero value:", z)
}
