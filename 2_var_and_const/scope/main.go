package main

import "fmt"

// scope refers to the visibility and
// accessibility of variables and constants

// in Go, there are three types of scope:
// package scope, function scope and block scope

// 1. package scope means that a variable or
// constant is accessible throughout the entire package

// 2. function scope means that a variable or constant
// is only accessible within the function where it is declared

// 3. block scope means that a variable or constant
// is only accessible within the block where it is declared
// like loops or {} block of code

const packageScopeConst = "I am a package scope constant"

func main() {
	functionScopeConst := "I am a function scope constant"

	fmt.Println(packageScopeConst) // accessible
	fmt.Println(functionScopeConst) // accessible

	{
		blockScopeConst := "I am a block scope constant"
		fmt.Println(blockScopeConst) // accessible
	}

	// fmt.Println(blockScopeConst) 
	// not accessible, will cause an error
}