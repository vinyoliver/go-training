package main

import (
	"fmt"
)

func main() {

	switch 4 {
	//the cases must be unique. They can't overlap
	case 1, 2, 3:
		fmt.Println("1...")
	case 4:
		fmt.Println("2...")
	default:
		fmt.Println("default...")
	}

	switch x := 1 + 3; x {
	case 1, 2, 3:
		fmt.Println("1, 2, 3...")
	case 4:
		fmt.Println("4...")
	default:
		fmt.Println("default...")
	}

	//the break is impplied;
	a := 10
	switch {
	case a > 10:
		fmt.Println("> 10")
	case a < 10:
		fmt.Println("< 10")
	default:
		fmt.Println("default A...")
		break
		fmt.Println("default B...")
	}

	//use fallthrough to keep goind after a case
	b := 21
	switch {
	case b > 10:
		fmt.Println(">> 10")
		fallthrough
	case b > 20:
		fmt.Println(">> 20")
	default:
		fmt.Println("default...")
	}

	//switch i.(type){....}

}
