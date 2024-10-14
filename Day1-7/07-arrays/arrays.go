package main

import "fmt"

func main() {
	// Declare an array of integers with a fixed size of 5
	var arr [5]int
	fmt.Println("Array:", arr) // Output: [0 0 0 0 0]

	// Initialize an array
	arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr) // Output: [1 2 3 4 5]

	// Access elements by index
	fmt.Println("First element:", arr[0]) // Output: 1

	// Modify an element by index
	arr[1] = 20
	fmt.Println("Modified Array:", arr) // Output: [1 20 3 4 5]

	// Get the length of the array
	fmt.Println("Array Length:", len(arr)) // Output: 5
}
