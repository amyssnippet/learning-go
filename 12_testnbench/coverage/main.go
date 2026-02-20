package main

// test coverage is a measure of how much of our code is executed
// when we run our tests. It helps us identify which parts of our code
// are not being tested and can be useful for improving our test suite.

// to check test coverage, we can use the `go test` command with the 
// `-cover` flag. This will show us the percentage of code that is covered
// by our tests. We can also generate a coverage report in HTML format
// using the `-coverprofile` flag and then view it in a browser.

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Run `go test -cover` to check test coverage")
}

// this can be run using `go test -cover` and `-coverprofile`
// in the terminal, and it will show 
// the percentage of code covered by the tests.

// and also visualize with `go tool cover -html` to identify 
// untested code paths and improve our test suite accordingly.

// this helps maintain quality standards and guide testing
// efforts for more relaiable and robust applications. 