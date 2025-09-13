package main

import (
	"sync"
	"testing"
)

// Benchmark with goroutines
func BenchmarkWithGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := Safecounter{v: make(map[string]int)}
		var wg sync.WaitGroup
		for j := 0; j < 1000; j++ {
			wg.Add(1)
			go func() {
				c.Inc("key")
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// Benchmark without goroutines
func BenchmarkWithoutGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := Safecounter{v: make(map[string]int)}
		for j := 0; j < 1000; j++ {
			c.Inc("key")
		}
	}
}
