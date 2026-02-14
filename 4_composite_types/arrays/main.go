package main

import "fmt"


func main() {
	// Arrays are fixed-size collections of elements of the same type.
	// The size of an array is determined at compile time and cannot be changed.
	// Arrays are value types, which means that when you assign an array to another variable, a copy of the array is created.

	// Declaring an array
	var arr [5]int // An array of 5 integers

	// Initializing an array
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	// You can also declare and initialize an array in one line
	arr2 := [5]int{1, 2, 3, 4, 5}

	// Accessing elements of an array
	fmt.Println(arr[0]) // Output: 1
	fmt.Println(arr2[1]) // Output: 2

	// The length of an array can be obtained using the built-in len function
	fmt.Println("Length of arr:", len(arr))
	fmt.Println("Length of arr2:", len(arr2))
}