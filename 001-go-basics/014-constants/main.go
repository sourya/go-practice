package main

import "fmt"

const a = 19
const b = 75.2
const c = "Russell James"
const d = "Ruby Genevieve"

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}
