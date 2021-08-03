package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/* os is a package and Open() is a method in it that
	   tries to open given files from the os*/
	file, err := os.Open("noName.txt")

	if err == nil {
		fmt.Println(file)
	} else {
		fmt.Println(err)
	}

	/* Create your own error object */
	myErrorObject := errors.New("My new error object")
	fmt.Println(myErrorObject)

	/* Handle error using ok*/
	attendance := map[string]bool{
		"Saumya":  true,
		"Ambika":  false,
		"Jyotsna": true}

	isPresent, ok := attendance["Saumya"]
	if ok {
		fmt.Println("Saumya is Present?", isPresent)
	} else {
		fmt.Println("No information about Saumya")
	}
}
