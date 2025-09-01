package main

import (
	"fmt"
	"time"
)

func main() {
	// Print "hello world" to the console
	fmt.Println("hello world")

	// Integer division example (result will be 2)
	fmt.Println(7 / 3)

	// Variables -- all variables must be used or needs to be removed
	// Explicit type declaration
	var name string = "example"
	fmt.Println(name)

	// Type inference - Go will infer the type
	var name2 = "example2"
	fmt.Println(name2)

	// Shorthand variable declaration (most common)
	name3 := "example3"
	fmt.Println(name3)

	// Constants
	const age int = 10
	fmt.Println(age)

	// Multiple constants declaration
	const (
		age2 = 15
		age3 = 20
	)
	fmt.Println(age2, age3)

	// Loop - Go only has for loops
	// While-loop style using for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// Infinite loop (commented out to prevent endless execution)
	// for {
	// 	println("1")
	// }

	// Traditional for loop with initializer, condition, and increment
	for i := 10; i <= 15; i++ {
		println(i)
	}

	// For loop with control statements
	for i := 10; i <= 15; i++ {
		if i == 12 {
			continue // Skip iteration when i equals 12
		}
		if i == 14 {
			break // Exit loop when i equals 14
		}
		println(i)
	}

	// Range-based for loop (iterates 0, 1, 2)
	for i := range 3 {
		println(i)
	}

	// Conditional statements
	var age4 = 18
	if age4 < 18 {
		println("Not an adult...")
	} else {
		println("Adult...")
	}

	// Else if ladder
	var age5 = 21
	if age5 < 18 {
		println("Not an adult...")
	} else if age5 < 25 {
		println("Young adult...")
	} else {
		println("Adult...")
	}

	// Multiple conditions with logical operators
	var role, hasPermission = "Admin", true
	if role == "Admin" || hasPermission {
		fmt.Println("Granted.")
	}

	// If statement with short variable declaration
	if age6 := 17; age6 < 18 {
		fmt.Println("Not an adult...")
	} else {
		fmt.Println("Adult...")
	}
	// Note: No ternary operator available in Go

	// Switch statement
	j := 3
	switch j {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}

	// Switch with multiple cases and time-based condition
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	case time.Friday:
		fmt.Println("End of Week")
	default:
		fmt.Println("Weekday")
	}

	// Type switch - determines the type of an interface{} value
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case float64:
			fmt.Println("Float")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other", t)
		}
	}

	// Test the type switch with different types
	whoAmI(1)
	whoAmI("golang")
	whoAmI(true)
	whoAmI(10.5)
	whoAmI(time.Now())
}
