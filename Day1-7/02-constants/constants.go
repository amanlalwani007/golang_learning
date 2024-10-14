package main

import "fmt"

func main(){
    // Declaring a constant
    const pi float64 = 3.14159
    fmt.Println(pi)

    // Declaring multiple constants at once
    const (
        x = 100
        y = "Hello"
    )
    fmt.Println(x, y)
}