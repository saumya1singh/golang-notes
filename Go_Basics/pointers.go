package main

import (
	"fmt"
)

func main() {
	/* pointers are the variables that stores the address of other variables
	   p pointer is pointing to an int variable
	*/
	var p *int

	if p != nil {
		fmt.Println("The value of p : ", *p)
	} else {
		fmt.Println("The value of p : nil")
	}

	var v int = 23
	p = &v
	if p != nil {
		fmt.Println("The value of p : ", *p)
	} else {
		fmt.Println("The value of p is nil")
	}

	var floatVariable float64 = 23.5
	pointer1 := &floatVariable
	fmt.Println("Value of pointer1 : ", *pointer1)

	/* lets change the value of pointer1
	 */
	*pointer1 = *pointer1 / 10
	fmt.Println("Updated value of pointer1 : ", *pointer1)
	fmt.Println("Value of floatVariable : ", floatVariable)
	/* both the values are same or simultaneously updated because of the line pointer1 := &floatVariable
	   it establishes common connection btw the pointer and variable */
}
