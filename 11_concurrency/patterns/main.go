package main

// there are concurrency patterns that can be used to 
// manage and coordinate goroutines effectively, such as:
// worker pools, fan-out/fan-in, pipelines, and publish/subscribe.

// these patterns help to structure concurrent code, 
// improve performance,
// and ensure proper synchronization and communication 
// between goroutines.

// 1. fan-out/fan-in: a pattern where multiple goroutines (fan-out)
// perform work concurrently and their results are collected 
// (fan-in) into a single channel for processing. eg:

func fanOutFanIn() {
	// create a channel to collect results
	results := make(chan int)

	// fan-out: start multiple goroutines to perform work
	for i := 0; i < 5; i++ {
		go func(n int) {
			// perform some work and send result to channel
			results <- n * n // example work: squaring the number
		}(i)
	}

	// fan-in: collect results from the channel
	for i := 0; i < 5; i++ {
		result := <-results
		println(result) // process the result (e.g., print it)
	}
}

// 2. pipelines: a pattern where data flows through a series of stages,
// with each stage running in its own goroutine and communicating 
// via channels. eg:

func pipeline() {
	// create channels for each stage of the pipeline
	stage1 := make(chan int)
	stage2 := make(chan int)

	// stage 1: generate numbers
	go func() {
		for i := 0; i < 5; i++ {
			stage1 <- i // send numbers to stage 1 channel
		}
		close(stage1) // close channel when done
	}()

	// stage 2: process numbers from stage 1
	go func() {
		for n := range stage1 { // read from stage 1 channel
			stage2 <- n * n // example processing: squaring the number
		}
		close(stage2) // close channel when done
	}()

	// collect results from stage 2
	for result := range stage2 {
		println(result) // process the result (e.g., print it)
	}
}

// 3. publish/subscribe: a pattern where publishers send messages to
// subscribers via channels, allowing for decoupled communication. eg:

func publishSubscribe() {
	// create a channel for publishing messages
	pub := make(chan string)

	// subscriber 1: listens for messages
	go func() {
		for msg := range pub {
			println("Subscriber 1 received:", msg)
		}
	}()

	// subscriber 2: listens for messages
	go func() {
		for msg := range pub {
			println("Subscriber 2 received:", msg)
		}
	}()

	// publisher: sends messages to the channel
	pub <- "Hello, Subscribers!"
	pub <- "Welcome to Go Concurrency Patterns"
	close(pub) // close channel when done
}

func main() {
	fanOutFanIn()
	pipeline()
	publishSubscribe()
}