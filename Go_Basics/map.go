/* map in go is an unordered collection of key value pair
its very common to use strings for keys and any other type for associated value*/

package main

import (
	"fmt"
)

func main() {
	/* creating a map called languages */
	languages := make(map[string]string)
	languages["one"] = "Java"
	languages["two"] = "C++"
	languages["three"] = "Python"
	languages["four"] = "Javascript"

	fmt.Println(languages)

	/* variables to print java */
	variableToPrintJava := languages["one"]
	fmt.Println(variableToPrintJava)

	/* delete a key value pair e.g key three */
	delete(languages, "three")
	fmt.Println(languages)

	/* Iterating over map key and value */
	fmt.Println("Iterating over map key and value pair")
	for k, v := range languages {
		fmt.Printf("%v : %v \n", k, v)
	}

}
