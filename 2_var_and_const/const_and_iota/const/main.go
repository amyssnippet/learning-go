package main

import "fmt"

// constants are immutable values
// that are known at compile time and
// do not change during the execution of the program.
// They are defined using the `const` keyword.

const name string = "Go Programming"

// multiple line constant declaration

const (
	pi float64 = 3.14159
	euler float64 = 2.71828
)

func main() {
	fmt.Println("Name(const string)", name)
	fmt.Println("Pi(const float64)", pi)
	fmt.Println("Euler's Number(const float64)", euler)
}