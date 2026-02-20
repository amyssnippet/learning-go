package main

import "fmt"

// variables are used to store data in memory
// multiple variables can be declared in a single line

var name, city string = "Amol", "Mumbai"

func main() {

	// var name type = value
	// this is a string variable named 'name' with the value "Amol"
	var name string = "Amol"

	fmt.Println("Name(string)", name)

	// this is a int variable named 'age' with the value 20
	var age int = 20

	fmt.Println("Age(int)", age)

	// this is a bool variable named 'isMarried' with the value false
	var isMarried bool = false

	fmt.Println("Is Married(bool)", isMarried)

	// this is a float64 variable named 'height' with the value 5.11
	var height float64 = 5.11

	fmt.Println("Height(float64)", height)

	// this is a float32 variable named 'weight' with the value 70.5
	var weight float32 = 70.5

	fmt.Println("Weight(float32)", weight)

	// the difference between float32 and float64 is the precision,
	// float64 can store more decimal places than float32


	// type inference, the data type is
	// inferred from the value assigned to the variable
	// to use the type inference, we can use the := operator
	city := "Mumbai"
	fmt.Println("City(inferred to string)", city)

	// the variable can be declared and 
	// assigned its value later
	var country string

	country = "India (before)"

	fmt.Println("Country(string)", country)

	// the variable can be reassigned a new value
	country = "India (after)"

	fmt.Println("Country(string)", country)
}