package main

// race detection is a tool that can be used to 
// detect race conditions in concurrent Go programs.
// It is enabled by running `go run -race main.go` 
// or `go test -race .` in the terminal.
// Race conditions occur when multiple goroutines 
// access shared data concurrently without proper synchronization.

func main() {
	// Example of a race condition
	var counter int

	// Increment the counter in multiple goroutines
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}

	// Wait for a moment to allow goroutines to finish
	// (not a proper way to wait, just for demonstration)
	select {}
}