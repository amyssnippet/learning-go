package main

// zero values are the default values assigned to variables
//  when they are declared without an explicit value
// the zero value for a string is an empty string ""
// the zero value for an int is 0
// the zero value for a bool is false
// the zero value for a float is 0.0
// the zero value for a pointer is nil

func main() {
	var name string
	var age int
	var isMarried bool
	var height float64
	var weight float32
	var pointer *int

	println("Name(string)", name)
	println("Age(int)", age)
	println("Is Married(bool)", isMarried)
	println("Height(float64)", height)
	println("Weight(float32)", weight)
	println("Pointer(*int)", pointer)
}