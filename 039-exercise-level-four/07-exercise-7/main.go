package main

import "fmt"

func main() {
	ab := []string{"Andrew", "Brudnak", "Ease of programming"}
	fmt.Println(ab)
	rb := []string{"Russell", "Brudnak", "With me you"}
	fmt.Println(rb)

	xp := [][]string{ab, rb}
	fmt.Println(xp)
	fmt.Printf("The xp variable type: %T\n", xp)

	for i, v := range xp {
		fmt.Println("record:", i)
		for j, value := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, value)
		}
	}
}
