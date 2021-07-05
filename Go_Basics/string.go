package main

import (
    "fmt"
    "strings"
)

func main() {

/*
string in Go can treated in 2 ways
1. as a simple single value
2. array of bytes
*/

    str1 := "This is implicitly typed"
    var str2 string = "This is explicitly typed"
    fmt.Println("Value -> ", str1)

    fmt.Printf("Value is : %v\nType is : %T\n", str1, str1)
    fmt.Printf("Value is : %v\nType is : %T\n", str2, str2)

    /* ToUpper and Title are the functions in package strings */
    fmt.Println(strings.ToUpper(str1))
    fmt.Println(strings.Title(str1))
    /* output of strings.ToUpper : THIS IS IMPLICITLY TYPED
       output of strings.Title : This Is Implicitly Typed
    */

    // check equality
    lValue := "hello"
    uValue := "HELLO"
    fmt.Println("Equal Case Sensitive?", lValue==uValue)
    fmt.Println("Equal Non-Case Sensitive", strings.EqualFold(lValue, uValue))

    // contains substring?
    fmt.Println("Contains substring?", strings.Contains(str1, "expli"))
    fmt.Println("Contains substring?", strings.Contains(str2, "expli"))


}


