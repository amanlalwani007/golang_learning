package main

import "fmt"

// A simple function that takes two integers and returns their sum
func add(x int, y int) int {
	return x + y
}

// A function that returns two values
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Function with named return values
func rectangleArea(width, height int) (area int) {
	area = width * height
	return // no need to explicitly return `area`
}

// Variadic function that sums any number of integers
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// A simple function
func greet(name string) string {
	return "Hello, " + name
}

func printMessage(f func(string) string, name string) {
	message := f(name)
	fmt.Println(message) // Output: Hello, Bob
}

// Closure example
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x // Closure captures and modifies the `sum` variable
		return sum
	}
}

func main() {
	result := add(5, 3)
	fmt.Println("Sum:", result) // Output: Sum: 8

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result) // Output: Result: 5
	}

	// Handling division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: cannot divide by zero
	}

	result = rectangleArea(5, 3)
	fmt.Println("Area:", result) // Output: Area: 15

	fmt.Println(sum(1, 2, 3, 4)) // Output: 10
	fmt.Println(sum(10, 20))     // Output: 30

	// Assign a function to a variable
	sayHello := greet
	fmt.Println(sayHello("Alice")) // Output: Hello, Alice

	// Pass a function as an argument to another function
	printMessage(sayHello, "Bob")

	// Anonymous function assigned to a variable
	add := func(a int, b int) int {
		return a + b
	}

	fmt.Println(add(5, 3)) // Output: 8

	adde := adder()
	fmt.Println(adde(10)) // Output: 10
	fmt.Println(adde(5))  // Output: 15
	fmt.Println(adde(3))  // Output: 18

}
