package main

import "fmt"

func main() {

	// Declare a map with string keys and int values
	var scores map[string]int
	scores = make(map[string]int) // Initialize the map

	// Add key-value pairs to the map
	scores["Alice"] = 95
	scores["Bob"] = 87

	// Access a value by key
	fmt.Println("Alice's score:", scores["Alice"]) // Output: 95

	// Modify a value
	scores["Alice"] = 98
	fmt.Println("Updated Alice's score:", scores["Alice"]) // Output: 98

	// Check if a key exists in the map
	score, exists := scores["Charlie"]
	if exists {
		fmt.Println("Charlie's score:", score)
	} else {
		fmt.Println("Charlie's score not found") // Output: Charlie's score not found
	}

	// Delete a key-value pair
	delete(scores, "Bob")
	fmt.Println("Scores:", scores) // Output: map[Alice:98]
	// Map literal with string keys and int values
	scores_with_literal := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}
	fmt.Println("Scores:", scores_with_literal) // Output: map[Alice:95 Bob:87 Charlie:92]

}
