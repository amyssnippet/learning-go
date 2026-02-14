package main

// tests in go are written in a separate file with the suffix _test.go
// and the test functions start with Test and take a pointer 
// to testing.T as an argument

// the files with _test.go suffix are not included in 
// the final build and are only used for testing purposes
// they can be run using the command `go test` in the terminal

import "testing"

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




// there are table driven tests which are a 
// common pattern in go testing
// they allow us to run the same test with 
// different inputs and expected outputs 
// and minimal code duplication

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Test 1", 2, 3, 5},
		{"Test 2", -1, 1, 0},
		{"Test 3", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Expected %d but got %d", tt.expected, result)
			}
		})
	}
}




// another thing is mocks and stubs which replace dependencies
// with controlled implementations for testing purposes

// stubs provide predefined responses to function calls 
// while mocks verify method calls and their parameters

// go's interface makes mocking natural essential for 
// tests without external dependencies

// there are also benchmarks in go which are used to 
// measure the performance of code and are written in 
// a similar way to tests but with the prefix Benchmark
// and take a pointer to testing.B as an argument

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, 3)
	}
}





// also there is a package named httptest which provides utilities
// for testing HTTP servers and clients without network calls
// it allows us to create test servers and clients for testing 
// HTTP handlers and clients in isolation

// the net/http/httptest package provides a way to create 
// test servers and clients for testing HTTP handlers and clients
// in isolation without making actual network calls

// we can use httptest.NewServer to create a test server and 
// httptest.NewRequest to create a test request for testing 
// HTTP handlers and clients

// example of using httptest to test an HTTP handler

/*
func TestHelloHandler(t *testing.T) {
	handler := http.HandlerFunc(helloHandler)

	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if string(body) != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!' but got '%s'", string(body))
	}
}
*/

// it is essential for testing HTTP handlers, middlewares and 
// http services