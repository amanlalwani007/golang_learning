package main

import "fmt"

func main() {
    // Basic for loop (similar to C-style loops)
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Simulating a while loop
    j := 0
    for j < 5 {
        fmt.Println(j)
        j++
    }

    // Infinite loop (useful when waiting for conditions, events, etc.)
    /*
    for {
        fmt.Println("Infinite loop")
    }
    */
}