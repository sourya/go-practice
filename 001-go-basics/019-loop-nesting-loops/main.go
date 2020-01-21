package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Printf("%03d Outter Loop\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t\t\t%v Inner Loop\n", j)
		}
	}
}
