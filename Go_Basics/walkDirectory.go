package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	/* this method Abs() returns 2 values
	   so we are keeping two values*/
	root, _ := filepath.Abs(".")
	fmt.Println("Visited Path :", root)

	/* this method walkFile returns error object */
	err := filepath.Walk(root, visitedPath)
	if err != nil {
		fmt.Println("error object is", err)
	}
}

/* this method returns error object and accept 3 param
3 params are - file path as string, info of the file and error object*/
func visitedPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory -", path)
		} else {
			fmt.Println("File -", path)
		}
	}
	return nil
}
