package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//create a reader object
	//bufio.NewReader will read the input from the input console
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text please: ")

	//a function with variable str that returns an error object. error object is presented by _
	// \n in single quotes means bytes not strings
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

}
