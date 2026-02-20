package main

// anonymous function is a function that does not have a name.
// It is defined using the func keyword without a name and can be
// assigned to a variable or passed as an argument to another function.

func main() {
	// defining an anonymous function and assigning it to a variable
	greet := func(name string) {
		println("Hello, " + name + "!")
	}

	// calling the anonymous function
	greet("Alice")

	// passing an anonymous function as an argument to another function
	execute(func() {
		println("This is an anonymous function passed as an argument.")
	})
}

// function that takes another function as an argument
func execute(f func()) {
	f()
}