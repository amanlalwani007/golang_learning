package main

import "fmt"

func main() {
	x := 10

	// Simple if-else condition
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is 5 or less")
	}

	// if with a variable declaration
	if y := 20; y > 15 {
		fmt.Println("y is greater than 15")
	} else {
		fmt.Println("y is 15 or less")
	}

	// else-if condition
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}
}