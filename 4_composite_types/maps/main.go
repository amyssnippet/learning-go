package main

// maps are unordered collections of key-value pairs.
// They are also known as hash maps or 
// dictionaries in other programming languages.
// Maps are reference types, which means that 
// when you assign a map to another variable, 
// both variables point to the same underlying data structure.

import "fmt"

func main() {
	// Declaring a map
	var m map[string]int // A map with string keys and int values

	// Initializing a map
	m = make(map[string]int) // Create an empty map using the make function

	// You can also declare and initialize a map in one line
	m2 := map[string]int{"one": 1, "two": 2, "three": 3}

	// Adding key-value pairs to a map
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	// Accessing values in a map using keys
	fmt.Println(m["one"]) // Output: 1
	fmt.Println(m2["two"]) // Output: 2

	// The length of a map can be obtained using the built-in len function
	fmt.Println("Length of m:", len(m))
	fmt.Println("Length of m2:", len(m2))

	// Deleting a key-value pair from a map
	delete(m, "two") // Remove the key "two" from the map
	fmt.Println("After deleting 'two' from m:", m)

	// Checking if a key exists in a map
	value, exists := m["three"]
	if exists {
		fmt.Println("Value for 'three':", value)
	} else {
		fmt.Println("'three' does not exist in the map")
	}


	// comma ok idiom is used to check if a key exists in a map
	value, ok := m["two"]
	if ok {
		fmt.Println("Value for 'two':", value)
	} else {
		fmt.Println("'two' does not exist in the map")
	}

	// Iterating over a map using a for loop
	for key, value := range m2 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}