package main

import (
	"fmt"
	"github.com/brudnak/go-practice/001-go-basics/097-example-tests/01-example/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	acdc.Sum(10, 10, 10)
}
