package main

// variadic function is a function that can take a variable 
// number of arguments. It is defined using the ... syntax 
// before the parameter type.

func main() {
	// calling the variadic function
	sum(1, 2, 3, 4, 5)
}

// variadic function definition
func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	println("The sum is:", total)
}


// multiple return values
func multipleReturns() (int, string) {
	return 42, "Hello"
}