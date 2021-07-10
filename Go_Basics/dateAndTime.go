package main

import (
	"fmt"
	"time"
)

func main() {

	/* Date and Now are members of the package time
	   you can't use them with objects created by you*/
	t := time.Date(2000, time.May, 26, 10, 45, 0, 0, time.UTC)
	fmt.Printf("I saw the world on %s\n", t)

	now := time.Now()
	fmt.Printf("Time now is : %s\n", now)

	/* Functions for the time object */
	fmt.Println("The month is : ", t.Month())
	fmt.Println("The Day is : ", t.Day())
	fmt.Println("The weekday is : ", t.Weekday())

	/* Add date OR
	   append date by one using AddDate Function*/
	nextDay := t.AddDate(0, 0, 1)
	fmt.Printf("Next Date After 26 May 2000 is : %v %v, %v (%v)\n",
		nextDay.Day(), nextDay.Month(), nextDay.Year(), nextDay.Weekday())

}
