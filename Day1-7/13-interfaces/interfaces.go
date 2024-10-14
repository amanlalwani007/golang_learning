package main

import "fmt"

// Declaring an interface
type Speaker interface {
	Speak() string
}

// A type that implements the Speaker interface
type Person struct {
	Name string
}

// Implementing the Speak method for Person
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

// Another type that implements the Speaker interface
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! I am " + d.Name
}

// Empty interface
func PrintAnything(value interface{}) {
	fmt.Println(value)
}

func main() {
	// Polymorphism: both types implement the Speaker interface
	var s Speaker

	p := Person{Name: "Alice"}
	d := Dog{Name: "Buddy"}

	s = p
	fmt.Println(s.Speak()) // Output: Hello, my name is Alice

	s = d
	fmt.Println(s.Speak()) // Output: Woof! I am Buddy

	PrintAnything("Hello")
	PrintAnything(42)
	PrintAnything(true)
}
