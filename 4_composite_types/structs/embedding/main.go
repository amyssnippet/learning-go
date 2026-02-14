package main

import "fmt"

// struct embedding includes one struct within another struct. 
// This allows the outer struct to access the fields and 
// methods of the embedded struct directly, 
// as if they were part of the outer struct.

func main() {
	type Person struct {
		Name string
		Age  int
	}

	type Employee struct {
		Person // Embedding the Person struct
		Position string
	}

	// Creating an instance of Employee
	emp := Employee{
		Person: Person{
			Name: "Alice",
			Age:  30,
		},
		Position: "Software Engineer",
	}

	// Accessing fields of the embedded struct directly
	fmt.Println("Name:", emp.Name)	   // Output: Name: Alice
	fmt.Println("Age:", emp.Age)         // Output: Age: 30
	fmt.Println("Position:", emp.Position) // Output: Position: Software Engineer
}