package main

import "fmt"



// Creating a custom type based on int
type Age int

// Define a custom type 'Age' with a method
func (a Age) IsAdult() bool {
	return a >= 18
}
func main() {
	var myAge Age = 30
	fmt.Println("My Age:", myAge) // Output: My Age: 30

	myAge = Age(20)
	fmt.Println("Is Adult?", myAge.IsAdult()) // Output: Is Adult? true

}
