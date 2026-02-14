package main

import "fmt"

// break statements are used to exit a loop prematurely.

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // Exit the loop when i is 5
		}
		fmt.Println(i) // Output: 1, 2, 3, 4
	}
}