package main

// generics is a powerful feature in Go 
// that allows you to write flexible and reusable code. 
// It enables you to create functions and data structures 
// that can work with any type, without sacrificing type safety. 
// This is achieved through the use of type parameters,
// which are specified using square brackets [].


// generic functions are defined using the func keyword,
// followed by the function name and a type parameter list in square brackets.
// The type parameters can be used within the function body to refer to the types that will be specified when the function is called.

// example of a generic function that takes a slice of any type and returns the first element:
func firstElement[T any](slice []T) T {
	return slice[0]
}


// generic types are defined using the type keyword,
// followed by the type name and a type parameter list in square brackets.
// The type parameters can be used within the type definition to refer to the types that will be specified when the type is instantiated.

// example of a generic type that represents a pair of values:
type Pair[T any] struct {
	First  T
	Second T
}

func main() {
	// using the generic function
	ints := []int{1, 2, 3}
	strings := []string{"a", "b", "c"}

	println(firstElement(ints))	// Output: 1
	println(firstElement(strings)) // Output: a
}




// generic interfaces are defined using the type keyword,
// followed by the interface name and a type parameter list in square brackets.
// The type parameters can be used within the interface definition to refer to the types that will be specified when the interface is implemented.

// example of a generic interface that represents a container for any type:
type Container[T any] interface {
	Get() T
	Set(value T)
}

// example of a struct that implements the Container interface for int type:
type IntContainer struct {
	value int
}

func (c *IntContainer) Get() int {
	return c.value
}

func (c *IntContainer) Set(value int) {
	c.value = value
}

// uncomment the main function below to see how to use the generic interface and the IntContainer struct.
// but comment the above main function to avoid duplicate main function error when running the code.

// func main() {
// 	// using the generic interface
// 	var c Container[int] = &IntContainer{}
// 	c.Set(42)
// 	println(c.Get()) // Output: 42
// }