package main

import "testing"

func BenchmarkCalc(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			calc("p096_sudoku.txt")
		}
	})
}
