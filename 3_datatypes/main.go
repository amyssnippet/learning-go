package main

// datatypes in Go are used to 
// specify the type of data that a variable can hold.
// Go has several built-in datatypes, including:
// - bool: represents a boolean value (true or false)
// - string: represents a sequence of characters
// - int: represents an integer value
// - float64: represents a floating-point number
// - complex128: represents a complex number

func main() {
	// declaring variables of different datatypes
	
	// there are 2 types of integers, signed and unsigned
	// signed integers can hold both positive and negative values
	// unsigned integers can only hold positive values

	// signed integers
	var age int = 30
	var temperature int = -5

	// unsigned integers
	var positiveAge uint = 30
	var positiveTemperature uint = 5

	// floating-point numbers
	var height float64 = 5.11
	var weight float32 = 70.5

	// boolean values
	var isMarried bool = false
	var isStudent bool = true

	// string values
	var name string = "Amol"
	var city string = "Mumbai"

	println("Age(int)", age)
	println("Temperature(int)", temperature)
	println("Positive Age(uint)", positiveAge)
	println("Positive Temperature(uint)", positiveTemperature)
	println("Height(float64)", height)
	println("Weight(float32)", weight)
	println("Is Married(bool)", isMarried)
	println("Is Student(bool)", isStudent)
	println("Name(string)", name)
	println("City(string)", city)

	// runes are used to represent Unicode characters
	// a rune is an alias for int32 and 
	// can hold any Unicode code point
	var letter rune = 'A'
	println("Letter(rune)", letter)


	// type conversion is the process of converting
	//  a value from one datatype to another
	var ageInYears int = 30
	var ageInMonths int = ageInYears * 12

	println("Age in Years(int)", ageInYears)
	println("Age in Months(int)", ageInMonths)

	// converting int to float64
	var heightInFeet float64 = float64(height) * 3.28084
	println("Height in Feet(float64)", heightInFeet)
}