package main

import "fmt"

// Passed by Value
// These types pass a copy of the value to the function. Modifications within the function do not affect the original data:

// Basic Types: int, float64, string, bool, etc.
// Arrays: [5]int, [3]string, etc.
// Structs: struct types (without using pointers)
// Function Pointers: The function signature itself is passed by value, but the pointed-to function operates on reference values
// Interfaces: The interface value itself is copied, but the actual implementation pointed to is a reference
// Passed by Reference
// These types are passed by reference, meaning that changes made within the function affect the original data:

// Pointers: *int, *struct, etc.
// Slices: []int, []string, etc.
// Maps: map[string]int, etc.
// Channels: chan int, etc.
// Function Values: Although technically passed by value, function values usually hold references to closures or other data that can be modified indirectly

func add(x int, y int) int {
	return x + y
}

//higher order function
func callback(x int, y int, myFunc func(int, int) int) int {
	defer fmt.Println("defered 3")
	fmt.Println("TESTING========")
	defer fmt.Println("defered 4")

	return myFunc(x, y)
}

//currying  example

// closures

func printValue(val int) {
	fmt.Println(val)
}

func main() {
	// defer fmt.Println("defered 2")

	fmt.Println(callback(10, 24, add))

	fmt.Println("Anomyous funcs :", callback(10, 24, func(i1, i2 int) int {
		return i1 * i2
	}))
	// defer fmt.Println("defered 1")

	x := 10
	defer printValue(x) // Argument 'x' is evaluated immediately
	x = 20
	printValue(x)
	// In Go, defer is a keyword that schedules a function call to be executed just before the surrounding function returns, however return value
	//  is set before the defer function call is executed (it means return values can be updated by defer function). only exception is the panic case , where defer gets executed first then the panic function is executed

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)

		// Output:
		// 2
		// 1
		// 0
	}

}
