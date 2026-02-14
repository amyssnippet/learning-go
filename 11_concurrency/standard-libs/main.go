package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// there are various standard libraries in Go that
// provide useful functions and types for different purposes.
// some commonly used standard libraries include:
// - fmt: for formatted I/O operations
// - math: for mathematical functions and constants
// - time: for working with dates and times
// - net/http: for building HTTP servers and clients
// - encoding/json: for encoding and decoding JSON data
// - sync: for synchronization primitives like mutexes and wait groups
// - io: for input and output operations
// - os: for interacting with the operating system
// - strings: for string manipulation functions
// - strconv: for converting strings to other types and vice versa

func fmtExample() {
	// Example of using the fmt package
	name := "Alice"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

func mathExample() {
	// Example of using the math package
	fmt.Println("Pi:", math.Pi)
	fmt.Println("Square root of 16:", math.Sqrt(16))
}

func timeExample() {
	// Example of using the time package
	now := time.Now()
	fmt.Println("Current time:", now)
}

func httpExample() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func jsonExample() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	person := Person{Name: "Bob", Age: 25}
	jsonData, _ := json.Marshal(person)
	fmt.Println("JSON:", string(jsonData))
}

func syncExample() {
	// Example of using the sync package
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
	}()

	wg.Wait()
}

func ioExample() {
	// Example of using the io package
	data := "Hello, World!"
	reader := strings.NewReader(data)
	buf := make([]byte, len(data))
	reader.Read(buf)
	fmt.Println("Read data:", string(buf))
}

func osExample() {
    // 1. Capture both the name and the error
    name, err := os.Hostname()

    // 2. Always check for errors in Go!
    if err != nil {
        fmt.Println("Error retrieving hostname:", err)
        return
    }

    // 3. Now you can use the single string value
    fmt.Println("Current hostname:", name)
}

func strconvExample() {
	// Example of using the strconv package
	numStr := "42"
	num, _ := strconv.Atoi(numStr)
	fmt.Println("Converted number:", num)
}

func main() {
	fmtExample()
	mathExample()
	timeExample()
	jsonExample()
	syncExample()
	ioExample()
	osExample()
	strconvExample()

	// Uncomment the following line to run the HTTP server example
	// httpExample()
}


// there are more and more standard libraries being added to Go,
// and they cover a wide range of functionalities.
// you can find the full list of standard libraries in the 
// official Go documentation: