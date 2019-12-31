package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Andrew","Last":"Brudnak","Age":29},{"First":"Russell","Last":"Brudnak","Age":2}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the data", people)

	for i, v := range people {
		i++
		fmt.Println("\nPERSON NUMBER:", i)
		fmt.Printf("First name: %v\t Last name: %v\t Age: %v\t", v.First, v.Last, v.Age)
		fmt.Println("")
	}
}
