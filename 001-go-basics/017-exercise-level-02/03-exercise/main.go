package main

import "fmt"

const myName = "andrew"
const myNum float64 = 3.14
const passGrade = 'A'

func main() {
	fmt.Println(myNum)
	fmt.Printf("%T\n", myNum)

	fmt.Println(myName)
	fmt.Printf("%T\n", myName)

	fmt.Println(passGrade)
	fmt.Printf("%T\n", passGrade)

	fmt.Println(string(passGrade))
}
