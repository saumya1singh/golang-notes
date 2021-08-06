package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from Go."
	/* create the File */
	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()

	/* Write something to the file using WriteString() method
	   it returns length of characters in the file */
	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with file of %v charters\n", length)
}

func checkError(err error) {
	if err != nil {
		/* Panic is a built-in function that stops the ordinary flow of control and begins panicking. */
		panic(err)
	}
}
