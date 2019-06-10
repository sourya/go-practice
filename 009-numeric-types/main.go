package main

import (
	"fmt"
	"runtime"
)

var age byte = 29

// byte = alias for uint8
// rune = alias for int32

// unsigned integers start at zero i.e. not negative

func main() {
	fmt.Println(age)
	fmt.Printf("%T\n", age)
	fmt.Println(runtime.GOOS)   // windows
	fmt.Println(runtime.GOARCH) // amd64
	// Dell Laptop
}
