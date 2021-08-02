package main

import (
	"fmt"
)

type Dog struct {
	Name   string
	Weight int
	Sound  string
}

/* methods of custom type e.g structure
method of type structure (dog)*/
/* these functions don't accept params, neither do they return any value */
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Ticy", 40, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()
}
