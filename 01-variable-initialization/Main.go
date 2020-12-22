package main

import (
	"fmt"
	"strconv"
)

//var block
var (
	name string = "Viny"
	age  int    = 31
)

func main() {
	var a int
	a = 24
	var b int32 = 45
  c := 33
  
	fmt.Printf("value of a: %v - type of %T\n", a, a)
	fmt.Printf("value of b: %v - type of %T\n", b, b)
	fmt.Printf("value of c: %v - type of %T\n", c, c)

	fmt.Println()

	//var scope use
	fmt.Printf("Hi may name is %v and I'm %v ", name, age)
	fmt.Println()

	x := 32
	fmt.Println(strconv.Itoa(x))
}
