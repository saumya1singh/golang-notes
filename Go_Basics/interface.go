package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

/* custom type Dog is implementation of interface Animal
since speak () is a method of type structure */

func (d Dog) Speak() string {
	return "Barkkk!"
}

func main() {
	/* wrapped the Dog object of structure around Interface Animal */
	poodle := Animal(Dog{"Teddy"})
	fmt.Println(poodle)
}
