package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("test i", i)
	}

	// x := 0
	// for ; x < 5; x++ {
	// 	fmt.Println("test x", x)
	// }

	// x := 0
	// for x < 5 {
	// 	fmt.Println("test x", x)
	// 	x++
	// }

	//do while...
	x := 0
	for {
		fmt.Println("test x", x)
		x++
		if x == 5 {
			break
		}
	}

	//we can loop over arrays, maps, and strings
	s := []int{1, 2, 3, 4, 5}
	for k, v := range s {
		fmt.Printf("%v - %v \n", k, v)
	}

	for k := range s {
		fmt.Printf("key %v \n", k)
	}

	for _, v := range s {
		fmt.Printf("value %v \n", v)
	}
}
