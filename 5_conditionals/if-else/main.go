package main

// if else statements are used to execute one block of code 
// if a specified condition is true, and another block if it is false.

func main() {
	age := 16

	if age >= 18 {
		println("You are an adult.") // This will not be executed
	} else {
		println("You are a minor.") // Output: You are a minor.
	}
}