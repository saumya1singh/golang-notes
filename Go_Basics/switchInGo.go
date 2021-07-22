package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/* the first line seeds the randomized value using a number of milliseconds */
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	result := ""

	switch dow {
	case 1:
		result = "It is sunday"
	case 7:
		result = "It is saturday"
	default:
		result = "It is weekday"
	}
	fmt.Println("Day of Week is : ", dow)
	fmt.Println("Result : ", result)

	x := -42
	switch {
	case x < 0:
		result = "x is less than zero"
	case x == 0:
		result = "x is equal to zero"
	default:
		result = "x is greater than zero"
	}
	fmt.Println(result)

}
