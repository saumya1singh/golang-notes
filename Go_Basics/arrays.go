package main

import (
	"fmt"
)

func main() {

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Orange"
	fmt.Println(colors)
	fmt.Println(colors[0])

	/* array declaration and assignment together */
	var numbers = [4]int{7, 23, 49, 34}
	fmt.Println(numbers)
	fmt.Println(numbers[3])
}
