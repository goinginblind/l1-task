package main

import (
	"testing"
)

// runs with [go test ./18 -bench=.]

// my results:
// goos: linux
// goarch: amd64
// pkg: github.com/goinginblind/l1-task/18
// cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx
// BenchmarkSafeCounterParallel-8     	15534568	        74.88 ns/op
// BenchmarkAtomicCounterParallel-8   	75641319	        14.99 ns/op
// PASS
// ok  	github.com/goinginblind/l1-task/18	2.402s

func BenchmarkSafeCounterParallel(b *testing.B) {
	var sc SafeCounter
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sc.Increment()
		}
	})
}

func BenchmarkAtomicCounterParallel(b *testing.B) {
	var ac AtomicCounter
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ac.Increment()
		}
	})
}
