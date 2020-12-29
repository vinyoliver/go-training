package main

import (
	"fmt"
)

type Doctor struct {
	name        string
	age         int
	specialties []string
}

func (d *Doctor) sayHello() {
	fmt.Printf("Hi my name is %v and I'm %v years old\n", d.name, d.age)
}

func (d *Doctor) GetOld() {
	d.age++
}

func (d *Doctor) hasSpecialitites() bool {
	if len(d.specialties) > 0 {
		return true
	} else {
		return false
	}
}

func main() {

	aDoctor := Doctor{
		age:         30,
		name:        "Jhon",
		specialties: []string{"a", "b", "c"},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.name)
	fmt.Println(aDoctor.specialties[1])

	aDoctor.sayHello()
	fmt.Println("Has Specialitites: ", aDoctor.hasSpecialitites())
	aDoctor.GetOld()
	aDoctor.sayHello()

	//anonymous struct
	person := struct {
		name string
		age  int
	}{name: "viny", age: 31}
	fmt.Println(person)

	//returning multiple values
	res, err := divide(3.0, 3.0)
	if err != nil {
		fmt.Println("Error division")
		return
	}

	fmt.Println(res)
}

//return multiple values
func divide(num1, num2 float64) (float64, error) {
	if num2 == 0.0 {
		return 0.0, fmt.Errorf("Error - division by zero")
	}
	return num1 / num2, nil
}
