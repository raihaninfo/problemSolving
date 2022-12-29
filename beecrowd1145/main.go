package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    for i := 1; i <= y; i++ {
        fmt.Print(i)
        if i % x == 0 {
            fmt.Println()
        } else {
            fmt.Print(" ")
        }
    }
}