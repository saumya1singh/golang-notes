/* new and make are two imp functions in terms of memory

new () function : allocate but doesn't initialize memory
make () function : allocate and initialize memory */

package main

import (
	"fmt"
)

func main() {

	/* declaration of objects e.g map without make function can cause panic
	example is below :
	*/

	var m map[string]int
	m["key"] = 23
	fmt.Println(m)

	/* above code snippet produce error
	   panic: assignment to entry in nil map
	*/

	/* if you intent to immediately add data to the map then declare
	it using make()*/
	m := make(map[string]int)
	m["key"] = 423
	fmt.Println(m)

	/* Memory de allocation
	memory is de allocated using GC - garbage collector
	 */
}
