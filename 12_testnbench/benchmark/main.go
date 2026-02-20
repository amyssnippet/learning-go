package main

// benchmarks measure code performance and are written
// in a similar way to tests but with the prefix Benchmark

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, 3)
	}
}

func add(a, b int) int {
	return a + b
}

// functions which start with Benchmark and use the 
// *testing.B type are recognized as benchmarks by the go test tool
// to test benchmarks, we run go test with the -bench 
// flag followed by a regex
// `go test -bench=.` will run all benchmarks in the package