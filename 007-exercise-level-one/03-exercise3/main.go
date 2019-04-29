package main

import "fmt"

var x = 19
var y = "Elon Musk"
var z = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
