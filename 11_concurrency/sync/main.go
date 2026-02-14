package main

import "sync"

// sync package provides basic synchronization primitives
// such as mutual exclusion locks. Other than the Once and
// WaitGroup types, most are intended for use by low-level
// library routines. Higher-level synchronization is better
// done with channels and communication.

// this package includes a Mutex type which
// provides a mutual exclusion lock.
// RWMutex is a reader/writer mutual exclusion lock.
// The lock can be held by an arbitrary

// WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines
// to wait for. Then each of the goroutines runs and calls
// Done when finished. At the same time, Wait can be used to
// block until all goroutines have finished.

// example of Mutex and RWMutex

func mutexExample() {
	// create a mutex
	var mu sync.Mutex

	// lock the mutex
	mu.Lock()

	// critical section of code that needs to be protected
	// ...

	// unlock the mutex
	mu.Unlock()
}

func rwMutexExample() {
	// create a RWMutex
	var rwMu sync.RWMutex

	// lock the RWMutex for writing
	rwMu.Lock()

	// critical section of code that needs to be protected for writing
	// ...

	// unlock the RWMutex for writing
	rwMu.Unlock()

	// lock the RWMutex for reading
	rwMu.RLock()

	// critical section of code that needs to be protected for reading
	// ...

	// unlock the RWMutex for reading
	rwMu.RUnlock()
}

// example of WaitGroup

func waitGroupExample() {
	var wg sync.WaitGroup

	// add the number of goroutines to wait for
	wg.Add(2)

	// start the first goroutine
	go func() {
		defer wg.Done() // signal that this goroutine is done
		// do some work
	}()

	// start the second goroutine
	go func() {
		defer wg.Done() // signal that this goroutine is done
		// do some work
	}()

	// wait for all goroutines to finish
	wg.Wait()
}

func main() {
	mutexExample()
	rwMutexExample()
	waitGroupExample()
}