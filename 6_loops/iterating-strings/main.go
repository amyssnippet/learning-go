package main

// iterating over strings in Go can be done using a for loop.
// When you iterate over a string, you get the index 
// and the rune (character) at that index.

import "fmt"

func main() {
	str := "Hello, Amol!"

	// Iterate over the string using a for loop
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}