package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// In Go, errors are values of the built-in type error
type error interface {
	Error() string
}

// creating custom error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Using fmt.Errorf for Formatted Error Messages:
func divide_format_error(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by %d", a, b)
	}
	return a / b, nil
}

// Defining a custom error type
type DivisionError struct {
	Dividend int
	Divisor  int
	Message  string
}

// Implementing the Error() method
func (e *DivisionError) Error() string {
	return fmt.Sprintf("Error: %s - Dividend: %d, Divisor: %d", e.Message, e.Dividend, e.Divisor)
}

// Function that returns a custom error
func divide_custom_error(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{
			Dividend: a,
			Divisor:  b,
			Message:  "division by zero",
		}
	}
	return a / b, nil
}

// Wrapping Errors
func someOperation() error {
	return errors.New("low-level error")
}

func anotherOperation() error {
	err := someOperation()
	if err != nil {
		return fmt.Errorf("anotherOperation failed: %w", err)
	}
	return nil
}

// panic and recover
func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()

	panic("something went wrong!") // This will be caught by recover
}
func main() {
	file, err := os.Open("non_existent_file.txt")
	if err != nil {
		// Handle the error
		fmt.Println("Error occurred:", err)

	}
	defer file.Close()

	fmt.Println("File opened successfully")
	// example of error checking

	i, err := strconv.Atoi("123a") // Error: invalid integer string
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Converted Integer:", i)

	// Creating custom errors
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: division by zero is not allowed

	}
	fmt.Println("Result:", result)

	// Using fmt.Errorf for Formatted Error Messages:
	result_format_error, err_format_error := divide_format_error(10, 0)
	if err_format_error != nil {
		fmt.Println("Error:", err_format_error) // Output: Error: cannot divide 10 by 0

	}
	fmt.Println("Result:", result_format_error)

	// Custom error types

	result_custom_error, err_custom_error := divide_custom_error(10, 0)
	if err_custom_error != nil {
		fmt.Println(err_custom_error) // Output: Error: division by zero - Dividend: 10, Divisor: 0

	}
	fmt.Println("Result:", result_custom_error)
	// wrapping Error
	err_wrapping := anotherOperation()
	if err_wrapping != nil {
		fmt.Println("Error:", err_wrapping) // Output: Error: anotherOperation failed: low-level error
	}

	// Unwrapping Error

	lowLevelErr := errors.New("low-level-error")
	wrappedErr := fmt.Errorf("high-level error: %w", lowLevelErr)
	// Unwrapping the error
	fmt.Println(errors.Unwrap(wrappedErr)) // Output: low-level error

	// Checking if an error matches
	if errors.Is(wrappedErr, lowLevelErr) {
		fmt.Println("The errors are the same") // Output: The errors are the same
	}

	// panic and Recover

	riskyOperation()
	fmt.Println("Program continues after panic")

}
