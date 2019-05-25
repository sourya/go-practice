package main

import "fmt"

func main() {
	m := map[string]int{
		"Andrew": 29,
		"Jessie": 29,
	}
	fmt.Println(m)

	delete(m, "Andrew")
	fmt.Println(m)

	delete(m, "Todd")
	fmt.Println(m)

	fmt.Println(m["Andrew"])
	fmt.Println(m["Jessie"])

	if v, ok := m["Jessie"]; ok {
		fmt.Println("value:", v)
		delete(m, "Jessie")
	}

	fmt.Println(m)
}
