package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {

    //onject to read input console
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter Number: ")
    str, _ := reader.ReadString('\n')

    f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Number is : ",f)
    }
}