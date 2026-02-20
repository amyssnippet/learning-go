package main

// iota is a predeclared identifier in Go 
// that represents successive untyped integer constants
// It is used in constant declarations 
// to simplify the definition of incrementing numbers.

const (
	// iota starts at 0 and increments by 1 for each constant in the block
	first = iota // first will be 0
	second        // second will be 1
	third         // third will be 2
)

// but what if i want to give iota to second
// and skip the first one, we can do that by assigning iota to second

const (
	_ = iota // skip the first value (0)
	first2    // first2 will be 1
	second2   // second2 will be 2
	third2    // third2 will be 3
)

func main() {
	println("First:", first)
	println("Second:", second)
	println("Third:", third)
	println("First2:", first2)
	println("Second2:", second2)
	println("Third2:", third2)
}