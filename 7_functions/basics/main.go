package main

// function is a block of code that performs a specific task.
// It is a reusable piece of code that can be called
// multiple times in a program. Functions can take 
// parameters and return values.

func main() {
	// calling the function
	greet()
}

// function definition
func greet() {
	println("Hello, World!")
}

// closure is a function that has access to variables 
// from its enclosing scope, even after the outer function 
// has finished executing.
func closure() func() {
	message := "Hello from the closure!"
	return func() {
		println(message)
	}
}

// named return values are return values that are 
// given a name in the function signature.
func namedReturnValues() (result int) {
	result = 42
	return
}

// call by value vs call by reference
// in Go, all function arguments are passed by value. 
// However, when you pass a pointer to a variable, 
// you can achieve the effect of call by reference.

func callByValue(x int) {
	x = 10
}

func callByReference(x *int) {
	*x = 10
}