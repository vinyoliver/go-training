package main

import (
	"fmt"
)

func main() {

	//if statements

	ages := map[string]int{
		"jhon": 35,
		"jane": 38,
		"viny": 31,
	}
	if found, ok := ages["jane"]; ok {
		fmt.Println("found: ", found)
	}

	values := []int{}
	if len(values) > 0 {
		fmt.Println("We have some values")
	} else {
		fmt.Println("We don't have any value")
	}

}
