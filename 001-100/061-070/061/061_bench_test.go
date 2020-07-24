package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func BenchmarkConcurrentCalc(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			calc(6, projecteuler.ConcurrentKind)
		}
	})
}

func BenchmarkRecursiveCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(6, projecteuler.RecursiveKind)
	}
}
