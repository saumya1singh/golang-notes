package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "./map.go"

	/* this function Readfile returns a byte array */
	content, err := ioutil.ReadFile(fileName)
	checkError(err)

	/* converting byte array into string */
	result := string(content)
	fmt.Println("Content of file: \n", result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
