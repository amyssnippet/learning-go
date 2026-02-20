package main

import "fmt"

// continue statements are used to skip the current 
// iteration of a loop and move to the next iteration.

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(i) // Output: 1, 3, 5, 7, 9
	}
}