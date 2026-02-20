package main

// switch statements are used to perform different actions 
// based on different conditions.

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		println("It's the start of the week.") // Output: It's the start of the week.
	case "Friday":
		println("It's almost the weekend.")
	default:
		println("It's a regular day.")
	}
}