package main

import "fmt"

func main() {
	// Declare a slice with values
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice) // Output: [1 2 3 4 5]

	// Slicing an array (slice is backed by an array)
	arr := [5]int{10, 20, 30, 40, 50}
	sliceFromArr := arr[1:4]                       // Get a slice from an array (indices 1 to 3)
	fmt.Println("Slice from Array:", sliceFromArr) // Output: [20 30 40]

	// Modifying the slice changes the underlying array
	sliceFromArr[0] = 25
	fmt.Println("Modified Slice:", sliceFromArr) // Output: [25 30 40]
	fmt.Println("Modified Array:", arr)          // Output: [10 25 30 40 50]

	// Append elements to a slice (automatically expands the slice)
	slice = append(slice, 6)
	fmt.Println("Appended Slice:", slice) // Output: [1 2 3 4 5 6]

	// Get the length and capacity of a slice
	fmt.Println("Length:", len(slice))   // Output: 6
	fmt.Println("Capacity:", cap(slice)) // Output: 6

    // Create a slice with make (length 3, capacity 5)
    slice_with_cap := make([]int, 3, 5)
    fmt.Println("Slice:", slice_with_cap)        // Output: [0 0 0]
    fmt.Println("Length:", len(slice_with_cap))  // Output: 3
    fmt.Println("Capacity:", cap(slice_with_cap)) // Output: 5

}