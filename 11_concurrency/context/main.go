package main

import (
	"context"
	"time"
)

// context package provides functions to manage deadlines,
// cancellation signals, and other request-scoped values across
// API boundaries and between processes.

// common uses of context include:
// http timeouts, database deadlines, goroutine cancellation,
// and passing request-scoped values.
// it is essential for webservers, microservices,circuit breakers
// and responsive apis that can handle timeouts and
// cancellations gracefully.

// examples of contexts:

func main() {
	ctx := context.Background() // create a base context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) // create a context with a timeout
	defer cancel() // ensure resources are cleaned up

	// use ctx in API calls, database queries, etc.
	// if the operation takes longer than 5 seconds, it will be cancelled
}