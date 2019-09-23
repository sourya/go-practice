package main

import "fmt"

const (
	name    string = "Andrew"
	year    int    = 1989
	married        = true
)

func main() {
	fmt.Printf("type of name: %T\t value: %v\n", name, name)
	fmt.Printf("type of year: %T\t\t value: %v\n", year, year)
	fmt.Printf("type of married: %T\t value: %v\n", married, married)
}
