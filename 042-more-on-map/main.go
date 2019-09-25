package main

import "fmt"

func main() {

	m := map[string]int{
		"Andrew":  29,
		"Jessica": 30,
		"Russell": 2,
		"Ruby":    1,
	}
	fmt.Println(m)
	fmt.Println(m["Russell"])
	fmt.Println(m["Ada"])

	v, ok := m["Ada"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Ada"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	if v, ok := m["Andrew"]; ok {
		fmt.Println("this should print an age:", v)
	}

	m["Penny"] = 10

	for i, v := range m {
		fmt.Printf("%v\"s age: %v\n", i, v)
	}
}
