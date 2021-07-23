package main

import (
	"fmt"
)

func main() {

	/* slice named colors */
	colors := []string{"red", "green", "blue"}
	fmt.Println(colors)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum : ", sum)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	fmt.Println("expecting same result as above")

	/* range will set the value of i to the current index */
	for i := range colors {
		fmt.Println(colors[i])
	}

	/* try break, continue and labels */
	sum = 1
	for sum < 1000 {

		sum += sum
		fmt.Println(sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break
		}
	}
endofprogram:
	fmt.Println("End of Program")
}
