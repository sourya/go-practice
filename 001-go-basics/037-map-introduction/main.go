package main

import "fmt"

func main() {
	m := map[string]int{
		"James":     3,
		"Genevieve": 1,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Genevieve"])
	fmt.Println(m["Ada"])

	v, ok := m["Roxanne"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["James"]; ok {
		fmt.Println(v)
		fmt.Println("in the IF")
		fmt.Println(ok)
	}
}
