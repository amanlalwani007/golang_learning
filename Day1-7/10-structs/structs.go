package main

import "fmt"

// Defining a struct called Person
type Person struct {
	Name string
	Age  int
}

func main() {
	// Creating an instance of the Person struct
	var person1 Person
	person1.Name = "Alice"
	person1.Age = 30
	fmt.Println("Person 1:", person1) // Output: {Alice 30}

	// Struct literal to create another instance
	person2 := Person{Name: "Bob", Age: 25}
	fmt.Println("Person 2:", person2) // Output: {Bob 25}

	// Accessing and modifying struct fields
	person2.Age = 26
	fmt.Println("Updated Person 2:", person2) // Output: {Bob 26}

	// Anonymous struct
	student := struct {
		Name  string
		Grade int
	}{
		Name:  "John",
		Grade: 10,
	}
	fmt.Println("Student:", student)

}
