package main

// shadowing is the ability to declare a new variable
//  with the same name as an existing variable

func main() {
	name := "Amol"
	println("Name:", name)

	// shadowing the variable name
	name = "John"
	println("Name:", name)

	// shadowing the variable name but in a deeper scope

	{
		name := "Alice"
		println("Name:", name)
	}

	println("Name:", name)
}