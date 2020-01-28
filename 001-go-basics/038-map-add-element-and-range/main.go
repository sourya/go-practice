package main

import "fmt"

func main() {

	mk := map[string]int{
		"Russell": 3,
		"Ruby": 1,
	}
	fmt.Println(mk)

	mk["Ada"] = 1

	for k, v := range mk {
		fmt.Println(k, v)
	}
	
}
