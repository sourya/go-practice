package main

import "fmt"

var x = 42
var y = "Russell & Ruby"
var z = true

func main() {
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
