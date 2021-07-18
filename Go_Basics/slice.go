package main

import (
	"fmt"
	"sort"
)

/* slice, very similar to array should have all data of same types
* but unlike array it is resizable */

func main() {
	/* size is not defined in slices unlike arrays */
	var colors = []string{"Blue", "Green", "Yellow"}
	fmt.Println(colors)

	/* append items in slice using append function */
	colors = append(colors, "Orange")
	fmt.Println(colors)

	/* remove First element of slice using append method */
	/* 	colors = append(colors[1:]) is same as 	colors = append(colors[1:len(colors])
	blank is by default taken as last*/
	colors = append(colors[1:])
	fmt.Println(colors)

	/* remove Last element of slice using append method */
	/* colors = append(colors[:len(colors)-1]]) is same as colors = append(colors[0:len(colors)-1]])
	blank is by default taken as the very first element*/
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	/* declare slice using make function and assign values later */
	numbers := make([]int, 5, 5)
	numbers[0] = 23
	numbers[1] = 7
	numbers[2] = 99
	numbers[3] = 89
	numbers[4] = 2
	fmt.Println(cap(numbers))

	numbers = append(numbers, 999)
	fmt.Println(numbers)

	/* print capacity of slice */
	fmt.Println(cap(numbers))

	/* sort slice in ascending order*/
	sort.Ints(numbers)
	fmt.Println(numbers)

}
