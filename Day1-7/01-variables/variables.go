package main

import "fmt"

func main() {
	// 1. Variables
	// Declare a variable with explicit type
	var x int = 10
	fmt.Println(x)

	// Declaring a variable without an explicit type (inferred)
	var y = 20
	fmt.Println(y)

	// Short variable declaration (type is inferred)
	z := 30
	fmt.Println(z)

	// Declare multiple variables at once
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
}
