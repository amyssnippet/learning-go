package main

// for loops are used to repeat a block of code a certain number
// of times or while a condition is true.

func main() {
	// Example of a for loop that counts from 1 to 5
	for i := 1; i <= 5; i++ {
		println(i) // Output: 1, 2, 3, 4, 5
	}

	// Example of a for loop that iterates over a slice
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
		// Output: 
		// Index: 0 Value: 10
		// Index: 1 Value: 20
		// Index: 2 Value: 30
		// Index: 3 Value: 40
		// Index: 4 Value: 50
	}
}