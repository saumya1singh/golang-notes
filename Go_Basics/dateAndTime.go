package main

import ("fmt"
        "time"
)

func main() {
    t := time.Date(2000, time.May, 26, 10, 45, 0, 0, time.UTC)
    fmt.Printf("I saw the world on %s\n", t)
}