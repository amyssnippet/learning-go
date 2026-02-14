package main

// go routines are lightweight threads managed by the Go runtime. 
// They allow you to run functions concurrently, 
// making it easier to perform tasks simultaneously 
// without blocking the main thread.
// it is used with the keyword "go" followed by a function call.

import (
	"fmt"
	"time"
)

func main() {
	go sayHello() // This will run concurrently with the main function
	fmt.Println("This is the main function.")
	time.Sleep(1 * time.Second) // Sleep to allow the goroutine to finish
}

func sayHello() {
	fmt.Println("Hello from the goroutine!")
}

// it is used for minimal memory overhead and efficient scheduling,
// making it run thousands and millions of goroutines 
// concurrently without significant performance degradation.

// it communicates using channels, 
// which are a powerful way to synchronize and 
// share data between goroutines.

// channels work on principle of share memory by communicating,
// where goroutines can send and receive values through channels, 
// allowing them to coordinate their actions and 
// share data safely without the need for explicit 
// locks or other synchronization mechanisms.

// typed conduits created with `make()`.
// channels come in two flavors: unbuffered and buffered.
// unbuffered channels require both sender and receiver 
// to be ready at the same time, 
// while buffered channels allow sending a 
// certain number of values without an immediate receiver.

// channels are used to synchronize goroutines and
// to pass data between them. 
// They provide a way to communicate and coordinate 
// the execution of goroutines, 
// making it easier to write concurrent programs in Go.

// example of using channels to communicate between goroutines:

func channelExample() {
	ch := make(chan string) // Create a channel of type string

	go func() {
		ch <- "Hello from the goroutine!" // Send a message to the channel
	}()

	message := <-ch // Receive the message from the channel
	fmt.Println(message) // Print the message
}

// In this example, we create a channel of type string and 
// start a goroutine that sends a message to the channel. 
// The main function then receives the message from the channel 
// and prints it. This demonstrates how channels can be used 
// to communicate between goroutines in Go.


// there is also a select statement that 
// allows you to wait on multiple channel operations,
// making it easier to handle multiple concurrent operations
// without blocking. 
// It allows you to specify multiple cases, 
// and the select statement will block until 
// one of the cases is ready to proceed.

func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

// In this example, we have two channels, ch1 and ch2, 
// and two goroutines that send messages 
// to these channels after a delay. 
// The select statement waits for either of 
// the channels to receive a message, 
// and once it does, it prints the message. 
// This allows us to handle multiple concurrent operations 
// without blocking, 
// as the select statement will proceed as soon 
// as one of the channels is ready.



// now if we look at the buffer vs unbuffered channels,
// unbuffered channels require both sender and 
// receiver to be ready at the same time, 
// while buffered channels allow sending a certain 
// number of values without an immediate receiver.

// example of unbuffered channel:

func unbufferedChannelExample() {
	ch := make(chan string) // Unbuffered channel

	go func() {
		ch <- "Hello from the goroutine!" // This will block until the main function receives the message
	}()

	message := <-ch // This will block until the goroutine sends the message
	fmt.Println(message)
}

// example of buffered channel:

func bufferedChannelExample() {
	ch := make(chan string, 2) // Buffered channel with capacity of 2

	ch <- "Message 1" // This will not block
	ch <- "Message 2" // This will not block

	go func() {
		fmt.Println(<-ch) // This will print "Message 1"
		fmt.Println(<-ch) // This will print "Message 2"
	}()

	time.Sleep(1 * time.Second) // Sleep to allow the goroutine to finish
}

// In the unbuffered channel example, 
// the sender will block until the receiver 
// is ready to receive the message. 
// In the buffered channel example,
// the sender can send messages without
// blocking until the buffer is full, 
// allowing for more flexibility in communication between goroutines.









// now another pattern is worker pools,
// which is a common concurrency pattern that allows
// you to manage a pool of fixed number of worker goroutines 
// to perform tasks concurrently. 
// This pattern is useful for efficiently handling 
// a large number of tasks without overwhelming the system.

// it controls resources usage while maintaining high concurrency,
// it allows you to limit the number of concurrent tasks,
// preventing resource exhaustion and improving performance.

// example of a worker pool:

func workerPoolExample() {
	tasks := make(chan int, 10) // Channel to send tasks
	results := make(chan int, 10) // Channel to receive results

	// Start worker goroutines
	for i := 0; i < 3; i++ {
		go worker(tasks, results)
	}

	// Send tasks to the workers
	for j := 0; j < 10; j++ {
		tasks <- j
	}
	close(tasks) // Close the tasks channel to signal no more tasks

	// Collect results from the workers
	for k := 0; k < 10; k++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}

func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- task * 2 // Process the task and send the result
	}
}

// In this example, we create a worker pool with 3 worker goroutines. 
// The main function sends tasks (integers) to the 
// workers through the tasks channel. 
// Each worker processes the task by multiplying it by 2
// and sends the result back through the results channel. 
// The main function collects and prints the results. 
// This pattern allows us to efficiently manage concurrent 
// tasks while controlling resource usage.