package main

import "fmt"

// Alias type declaration
type ID = string

func main() {
	var userID ID = "ABC123"
	fmt.Println("User ID:", userID) // Output: User ID: ABC123
}

// Type Alias: Aliased types are completely interchangeable with the original type. There is no distinction.
// Custom Types: Custom types are treated as new, distinct types and aren't directly interchangeable with the original type,
//  even if based on the same underlying type.
