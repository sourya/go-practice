package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person
type ByName []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age > a[j].Age }

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := Person{"Andrew", 29}
	p2 := Person{"Jessica", 30}
	p3 := Person{"Russell", 2}
	p4 := Person{"Ruby", 1}

	people := []Person{p4, p2, p3, p1}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println()
	fmt.Println("----------")
	fmt.Println()

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
