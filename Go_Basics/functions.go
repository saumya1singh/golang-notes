package main

import (
	"fmt"
)

func main() {
	firstFunction()
	sum := 0
	sum = functionWithParam(10, 90)
	fmt.Println("Sum : ", sum)

	sum = addAllValues(90, 50, 10, 20)
	fmt.Println("New Sum : ", sum)
}

func firstFunction() {
	fmt.Println("First Function in Go")
}

/* function with argument and return type int */
func functionWithParam(value1 int, value2 int) int {
	return value1 + value2
}

/*You can create functions that expect arbitrary number of values
as long as they all are of same types */
func addAllValues(values ...int) int {
	sum := 0

	/* add all values like you do in slice */
	for i := range values {
		sum += values[i]
	}
	/* print type of slice */
	fmt.Printf("%T\n", values)
	return sum
}
