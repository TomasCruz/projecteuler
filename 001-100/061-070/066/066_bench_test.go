package main

import (
	"testing"
)

func BenchmarkChakravalaCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(1000, chakravalaSolver{})
	}
}

func BenchmarkContinuedFractionCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc(1000, makeContinuedFractionSolver(1000))
	}
}
