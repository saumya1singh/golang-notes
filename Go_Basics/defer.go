package main

import (
	"fmt"
)

/* defer as the words is self-explanatory
wherever any statement is marked as defer that is postponed
*/

/* within a function if any statement is marked as defer then that is shown by the end of function
when all other statements are already executed*/
func main() {
	defer fmt.Println("Defer 1 - It is a deferred statement- will be printed in last")
	fmt.Println("Not deferred statement")
	myFunc()
	defer fmt.Println("Defer 2 - It is a deferred statement- will be printed in second last")
	fmt.Println("Defer statements are printed in LIFO manner like stack. \nSo last deferred statement in a function willl be shown first")

}

func myFunc() {
	fmt.Println("I will be executed after 1st non deferred statement of main method")
}
