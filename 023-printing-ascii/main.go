package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Println("value of i: ", i)
		fmt.Printf("\t\tdecimal: %d\t hex: %#x\t unicode: %#U\n", i, i, i)
	}
}
