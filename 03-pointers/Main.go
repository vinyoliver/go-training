package main

import (
	"fmt"
)

func main() {
	//when you prefix a type with * you get a pointer
	//a pointer that has no value is nil

	var a int = 10
	var b *int = &a
	c := &a

	// *variable gets the data
	// &variable gets the memory reference
	fmt.Println(&a, *b)
	fmt.Println("Value of b", *b)
	fmt.Println("Reference of b", b)
	fmt.Println("Value of c", c)

	printValues := func() {
		fmt.Printf("a is %d, b is %d, c is %d \n", a, *b, *c)
		fmt.Printf("a refs %v, b refs %v, c refs %v", &a, b, c)
	}

	a = 20
	printValues()
}
