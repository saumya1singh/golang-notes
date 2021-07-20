/* map in go is an unordered collection of key value pair
its very common to use strings for keys and any other type for associated value*/

package main

import (
	"fmt"
	"sort"
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

	/* Sorting maps result using SLICES */
	/* when iterating iver range keyword if you only assign one variable you
	will only get the key and not value*/

	keys := make([]string, len(languages))
	i := 0
	for k := range languages {
		keys[i] = k
		i++
	}

	/* sort the slice named keys alphabetically */
	sort.Strings(keys)
	fmt.Println("Print the slice", keys)
	//     fmt.Println("Now print sorted map")
    for i := range languages {
        fmt.Printf("%v : %v\n", i, languages[i])
    }
}
