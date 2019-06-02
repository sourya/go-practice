package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 2, 4, 3, 1}
	xs := []string{"Eagle", "Baboon", "Aardvark", "Dolphin", "Cheetah"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println("")
	fmt.Println(xi)
	fmt.Println(xs)
}
