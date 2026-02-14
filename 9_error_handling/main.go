package main

import (
	"errors"
	"fmt"
	"os"
)

// error handling in Go is done using the built-in error type,
// which is a simple interface that has a single method,
// Error(), that returns a string.

// example of a function that returns an error:
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}


// there are patterns to handle errors in Go,
// such as using the "comma ok" idiom to check for the 
// presence of a value in a map, or using the "defer" 
// keyword to ensure that resources are released even 
// if an error occurs.

// example of using the "comma ok" idiom to check for 
// the presence of a value in a map:

func commaOkExample() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	value, ok := m["c"]
	if !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println("Value:", value)
	}
}

// example of using the "defer" keyword to ensure that 
// resources are released even if an error occurs:

func deferExample() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// do something with the file
}



// one more pattern is if err!= nil {
// 	// handle the error
// 	return
// }
// this pattern is used to check for errors and return 
// early if an error occurs,
// which helps to keep the code clean and easy to read.

// example of using the early return pattern to handle errors:

func earlyReturnExample() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}





// to print errors, there are methods like errors.New() and 
// fmt.Errorf() that can be used to create error values 
// with custom messages.

// example for errors.New() and fmt.Errorf():

func errorsNewExample() {
	err1 := errors.New("this is an error from errors.New()")
	err2 := fmt.Errorf("this is an error with a custom message")
	fmt.Println("Error1:", err1)
	fmt.Println("Error2:", err2)
}

func packageExamples() {
	errorsNewExample()
}





// errors can be wrapper and unwrapped using the errors package,
// which provides functions like errors.Wrap() and errors.Unwrap()
// to add context to errors and retrieve the original error.

// example of wrapping and unwrapping errors:

func wrapExample() {
	err := errors.New("original error")
	wrappedErr := fmt.Errorf("wrapped error: %w", err)
	fmt.Println("Wrapped Error:", wrappedErr)

	unwrappedErr := errors.Unwrap(wrappedErr)
	fmt.Println("Unwrapped Error:", unwrappedErr)
}



// sentinel errors are predefined error values that can be used to
// represent specific error conditions, such as io.EOF for end of 
// file errors, or os.ErrNotExist for file not found errors.

// example of using sentinel errors:

func sentinelExample() {
	_, err := os.Open("nonexistentfile.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File not found")
	} else if err != nil {
		fmt.Println("Error opening file:", err)
	} else {
		fmt.Println("File opened successfully")
	}
}





// there are methods like panic() and recover() that can be 
// used to handle unexpected errors

// panic() is used to raise a runtime error, 
// which will stop the normal execution of the program 
// and print a stack trace.

// example of using panic():

func panicExample() {
	panic("something went wrong")
}

// recover() is used to catch a panic and prevent the program 
// from crashing, allowing you to handle the error gracefully.

// example of using recover():

func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something went wrong")
}







// Go automatically prints stack traces on panic showing 
// call chain and line numbers, which can be helpful for debugging.

// example of automatic stack trace on panic:

func automaticStackTraceExample() {
	panic("something went wrong")
}