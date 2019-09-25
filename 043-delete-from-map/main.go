package main

import "fmt"

func main() {
	m := map[int]string{
		1: "Andrew",
		2: "Lewis",
		3: "Ivan",
	}

	fmt.Println(m)

	delete(m, 1)

	if v, ok := m[2]; ok {
		fmt.Println("value:", v)
		delete(m, 2)
	}

	fmt.Println(m)
	fmt.Println(m[2])
	fmt.Println(m[3])
}
