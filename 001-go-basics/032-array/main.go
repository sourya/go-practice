package main

import "fmt"

func main() {

	// Arrays are useful when planning the detailed layout of memory
	// and sometimes can help avoid allocation, but primarily they are
	// a building block for slices, the subject of the next section. To
	// lay the foundation for that topic, here are a few words about arrays.

	// There are major differences between the ways arrays work in Go and C. In Go,

	// Arrays are values. Assigning one array to another copies all the elements.
	// In particular, if you pass an array to a function, it will receive a copy of
	// the array, not a pointer to it.
	// The size of an array is part of its type. The types [10]int and [20]int are distinct.
	// The value property can be useful but also expensive; if you want C-like behavior and
	// efficiency, you can pass a pointer to the array.

	var x [5]int
	fmt.Println(x)
	x[3] = 19
	fmt.Println(x)
	fmt.Println(len(x))
}
