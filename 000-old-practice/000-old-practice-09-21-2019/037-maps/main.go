package main

import "fmt"

func main() {
	m := map[string]int{
		"Andrew":  29,
		"Jessie":  29,
		"Russell": 2,
		"Ruby":    1,
	}
	fmt.Println(m)
	fmt.Println(m["Russell"])

	fmt.Println(m["Shelby"])

	v, ok := m["Shelby"]

	fmt.Println(v)
	fmt.Println(ok)

	m["Todd"] = 33

	if v, ok := m["Jessie"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// range over slice
	xi := []int{1, 2, 3, 4, 5, 6}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
