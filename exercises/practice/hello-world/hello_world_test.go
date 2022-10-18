package hash

import "testing"

// Define a function named Hash that takes a string,
// and returns a uint32.
// In other words, define a function with the following signature:
// Hash() uint32

func TestHash(t *testing.T) {
	val := "Hello, World!"
	expected := uint32(1525479220)
	if observed := Hash(val); observed != expected {
		t.Fatalf("Hash(val) = %v, want %v", observed, expected)
	}
}

// BenchmarkHash() is a benchmarking function. These functions follow the
// form `func BenchmarkXxx(*testing.B)` and can be used to test the performance
// of your implementation. They may not be present in every exercise, but when
// they are you can run them by including the `-bench` flag with the `go test`
// command, like so: `go test -v --bench . --benchmem`
//
// You will see output similar to the following:
//
// BenchmarkHash   	2000000000	         0.46 ns/op
//
// This means that the loop ran 2000000000 times at a speed of 0.46 ns per loop.
//
// While benchmarking can be useful to compare different iterations of the same
// exercise, keep in mind that others will run the same benchmarks on different
// machines, with different specs, so the results from these benchmark tests may
// vary.
func BenchmarkHash(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Hash("Hello, World!")
	}
}
