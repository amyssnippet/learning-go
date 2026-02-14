package main

// for range loops are used to iterate over elements 
// in a collection, such as arrays, slices, maps, or strings.

func main() {
	// Example of a for range loop that iterates over a slice
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

	// Example of a for range loop that iterates over a map
	person := map[string]string{
		"name": "Alice",
		"age":  "30",
	}
	for key, value := range person {
		println("Key:", key, "Value:", value)
		// Output:
		// Key: name Value: Alice
		// Key: age Value: 30
	}

	// Example of a for range loop that iterates over a string
	str := "Hello"
	for index, char := range str {
		println("Index:", index, "Character:", string(char))
		// Output:
		// Index: 0 Character: H
		// Index: 1 Character: e
		// Index: 2 Character: l
		// Index: 3 Character: l
		// Index: 4 Character: o
	}
}