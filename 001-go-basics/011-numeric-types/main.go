package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

func main() {
	x = 75
	y = 3.14
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
