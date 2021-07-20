package main

import (
	"fmt"
)

/* struct is a data structure in go
similar to C's struct and Java classes

It doesn't have inheritance model like super or substruct etc */

type Flower struct {
	Name        string
	Color       string
	PetalNumber int
}

func main() {

	sunFlower := Flower{"SunFlower", "Yellow", 50}
	rose := Flower{"Rose", "Red", 80}
	fmt.Println(sunFlower)
	fmt.Println(rose)

	/* detailed info about the structure object */
	fmt.Printf("%+v\n", sunFlower)

	/* separately get the objects feature */
	fmt.Printf("Name : %v, Color : %v\n", rose.Name, rose.Color)
}
