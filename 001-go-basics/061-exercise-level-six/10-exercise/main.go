package main

import "fmt"

func main() {

	foo()
	fmt.Println("Next line will try printing variable 'x' from foo in main")
	// line below WON'T run
	// variable 'x' is SCOPED to func foo
	// fmt.Print(x)
	x := "Andrew"
	fmt.Println("Declared 'x' inside main, it's value is:", x)
	foo()

	b := bar()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func foo() {
	x := 19
	fmt.Println("Variable 'x' declared inside func foo:", x)
}

// increment - great simple example of closure
func bar() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
