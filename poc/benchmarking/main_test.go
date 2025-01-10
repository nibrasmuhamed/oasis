package main

import "testing"

func BenchmarkRandomGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randomGenerator(100) // You can adjust the argument as needed
	}
}
