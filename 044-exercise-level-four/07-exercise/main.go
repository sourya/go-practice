package main

import "fmt"

func main() {
	ab := []string{"Andrew", "Brudnak", "Scratch, Sniff, Wiggle, Wobble"}
	rb := []string{"Russell", "James", "Bring my elephant home!"}
	xp := [][]string{ab, rb}

	for i, v := range xp {
		fmt.Println("Record:", i+1)
		for j, val := range v {
			fmt.Printf("\t\tIndex postion: %v \t Value: %v \n", j, val)
		}
	}
}
