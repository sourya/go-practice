package main

import "fmt"

func main() {
	foo()
	bar("Ruby")
	s1 := woo("Russell")
	fmt.Println(s1)
	x, y := mouse("Andrew", "Brudnak")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identifier(parameters) (returns(s)) { ... }
func foo() {
	fmt.Println("Hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ", says hello")
	b := true
	return a, b
}
