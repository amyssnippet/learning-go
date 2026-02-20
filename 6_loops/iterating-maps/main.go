package main

// iterating over maps in Go can be done using the range keyword.
// When you iterate over a map, you get two values: 
// the key and the value.

import "fmt"

func main() {
	// Create a map of string keys and int values
	m := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	// Iterate over the map using a for range loop
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}