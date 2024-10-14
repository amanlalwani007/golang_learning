package main

import "fmt"

func modifyValue(val int) {

	val = 100
}

func modifyPointer(ptr *int) {
	*ptr = 100
}

type Person struct {
	name string
	age  int
}

func updateAge(p *Person, newAge int) {
	p.age = newAge
}

func main() {

	x := 10
	p := &x // p is a pointer to x

	fmt.Println("Value of x:", x)   // Output: 10
	fmt.Println("Address of x:", p) // Output: memory address of x

	// Dereferencing the pointer
	fmt.Println("Value at pointer p:", *p) // Output: 10

	// Modifying the value using the pointer
	*p = 20
	fmt.Println("New value of x:", x) // Output: 20

	// In Go, all function arguments are passed by value by default. This means that when a variable is passed to a function,
	// a copy of that variable is made. Modifying the copy inside the function does not affect the original variable.
	y := 10
	modifyValue(y)                                    // Passing x by value
	fmt.Println("Value of y after function call:", y) // Output: 10 (unchanged)

	// Although Go does not have pass-by-reference in the traditional sense, you can achieve similar behavior by passing a
	// pointer to a variable. When you pass a pointer, the function can modify the original variable’s value because it has
	// access to its memory address.

	z := 10
	modifyPointer(&z)                                 // Passing a pointer to x
	fmt.Println("Value of z after function call:", z) // Output: 100 (modified)

	// Pointers in struct

	person := Person{name: "Alice", age: 30}
	updateAge(&person, 35)                  // Passing a pointer to the struct
	fmt.Println("Updated age:", person.age) // Output: 35

	// pointers and Nil

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is nil")
	}

}
