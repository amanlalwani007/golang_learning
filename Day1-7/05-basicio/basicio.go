package main

import "fmt"

func main() {
	var name string
	var age int

	// Taking input from the user
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // Scanln reads input until a newline

	fmt.Print("Enter your age: ")
	fmt.Scan(&age) // Scan reads input without a newline

	fmt.Printf("Your name is %s and you are %d years old\n", name, age)
}
