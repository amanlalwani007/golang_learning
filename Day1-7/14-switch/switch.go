package main

import "fmt"

func main() {

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday":
		fmt.Println("It's Tuesday") // Output: It's Tuesday
	case "Friday":
		fmt.Println("Almost the weekend!")
	default:
		fmt.Println("It's a regular day")
	}

	x := 10

	switch {
	case x < 5:
		fmt.Println("x is less than 5")
	case x == 10:
		fmt.Println("x is exactly 10") // Output: x is exactly 10
	default:
		fmt.Println("x is greater than 5 but not 10")
	}

	day_multiple_switch := "Saturday"

	switch day_multiple_switch {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!") // Output: It's the weekend!
	default:
		fmt.Println("It's a weekday")
	}

	// fallthrough in switch
	day_fallthrough := "Monday"

	switch day_fallthrough {
	case "Monday":
		fmt.Println("It's Monday")
		fallthrough
	case "Tuesday":
		fmt.Println("It's Tuesday") // Output: It's Tuesday
	default:
		fmt.Println("Default case")
	}

}
