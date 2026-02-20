package main

// struct tags provide metadata about the fields of a struct.
// They are used to specify how the fields should be encoded
// or decoded when working with data formats like JSON, XML, etc.

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"` // The json tag specifies the key name for JSON encoding/decoding
	Age  int    `json:"age"`
}

func main() {
	// Creating an instance of Person
	p := Person{Name: "Alice", Age: 30}

	// Encoding the struct to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData)) // Output: JSON data: {"name":"Alice","age":30}

	// Decoding JSON back to a struct
	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Decoded struct: %+v\n", p2) // Output: Decoded struct: {Name:Alice Age:30}
}