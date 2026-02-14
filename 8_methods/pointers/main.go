package main

import "fmt"

// pointers allow you to work with memory addresses directly.
// They are useful for optimizing performance and managing
// resources efficiently. In Go, you can declare a pointer
// using the * operator, and you can get the address of a
// variable using the & operator.

func main() {
	// Example of using pointers
	var x int = 10
	var p *int = &x // p is a pointer to x

	fmt.Println("Value of x:", x)	   // Output: 10
	fmt.Println("Address of x:", &x)    // Output: memory address of x
	fmt.Println("Value of p:", p)       // Output: memory address of x
	fmt.Println("Value at p:", *p)     // Output: 10

	// Modifying x through the pointer
	*p = 20
	fmt.Println("New value of x:", x)   // Output: 20
}