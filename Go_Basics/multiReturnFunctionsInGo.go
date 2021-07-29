package main

import (
	"fmt"
)

func main() {
	n1, l1 := FullName("Helen", "Keller")
	fmt.Printf("Full Name is %v , total length is %v \n", n1, l1)

	n2, l2 := FullNameNakedReturn("B.K.", "Shivani")
	fmt.Printf("Full Name is %v , total length is %v \n", n2, l2)
}

/* order in which params are written in func signature should
be same as at time of return in end */
func FullName(f, l string) (string, int) {
	fullName := f + " " + l
	length := len(fullName)
	return fullName, length
}

/* This way is called naked Return in which after return statement nothing else us needed  */
func FullNameNakedReturn(f, l string) (fullName string, length int) {
	fullName = f + " " + l
	length = len(fullName)
	return
}
