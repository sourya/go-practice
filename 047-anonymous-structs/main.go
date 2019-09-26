package main

import "fmt"

//type person struct {
//	first string
//	last string
//}

func main() {

	// anonymous struct!
	// why?
	// limit code pollution
	p1 := struct {
		first string
		last  string
	}{
		first: "Ruby",
		last:  "Brudnak",
	}

	fmt.Println(p1)
}
