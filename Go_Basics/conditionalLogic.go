package main

import (
	"fmt"
)

/* Go is sensitive to whitespaces
put else in same line as ending braces of if statement

if-else is same as java etc except that parenthesis is not needed*/

func main() {
	// 	var x float64 = 43
	var result string

	if x := 43; x < 0 {
		result = "x is smaller than 0"
	} else if x == 0 {
		result = "x is equal to 0"
	} else {
		result = "x is greater than 0"
	}

	fmt.Println("Result is : ", result)
}
