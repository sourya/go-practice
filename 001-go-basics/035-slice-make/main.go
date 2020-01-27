package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)

	ms := make([]int, 10, 12)
	fmt.Println(ms)
	fmt.Println(len(ms))
	fmt.Println(cap(ms))

	ms[0] = 19
	ms[9] = 999
	ms[4] = 99999999
	fmt.Println(ms)

	ms = append(ms, 5454)
	ms = append(ms, 5455)
	ms = append(ms, 5456)

	fmt.Println(len(ms))
	fmt.Println(cap(ms))
}
