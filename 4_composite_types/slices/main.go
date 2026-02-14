package main

// slices are dynamic-size collections of elements of the same type.
// They are built on top of arrays and 
// provide more flexibility in terms of size and functionality. 
// Slices are reference types, which means that 
// when you assign a slice to another variable, 
// both variables point to the same underlying array.

import "fmt"

func main() {
	// Declaring a slice
	var s []int // A slice of integers

	// Initializing a slice
	s = []int{1, 2, 3, 4, 5}

	// You can also declare and initialize a slice in one line
	s2 := []int{6, 7, 8, 9, 10}

	// Accessing elements of a slice
	fmt.Println(s[0]) // Output: 1
	fmt.Println(s2[1]) // Output: 7

	// The length of a slice can be obtained using the built-in len function
	fmt.Println("Length of s:", len(s))
	fmt.Println("Length of s2:", len(s2))

	// Slices can be resized using the built-in append function
	s = append(s, 6) // Append an element to the slice
	fmt.Println("After appending to s:", s)

	// Slices can also be created from arrays
	// array to slice conversion
	arr := [5]int{1, 2, 3, 4, 5}
	s3 := arr[1:4] // Create a slice from the array (elements at index 1 to 3)
	fmt.Println("Slice s3 from array:", s3)

	// Modifying a slice will affect the underlying array
	s3[0] = 10
	fmt.Println("Modified slice s3:", s3)
	fmt.Println("Underlying array after modifying s3:", arr)

	// Slices can also be created using the make function
	s4 := make([]int, 5) // Create a slice of integers with length 5
	fmt.Println("Slice s4 created with make:", s4)

	// slice to array conversion
	arr2 := [5]int{1, 2, 3, 4, 5}
	s5 := arr2[:] // Create a slice from the entire array
	fmt.Println("Slice s5 from array:", s5)
}