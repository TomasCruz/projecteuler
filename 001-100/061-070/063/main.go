package main

import (
	"math"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 63; Powerful digit counts

The 5-digit number, 16807=7^5, is also a fifth power. Similarly, the 9-digit number,
134217728=8^9, is a ninth power. How many n-digit positive integers exist which are also an nth power?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	// a = x^n, 10^n-1 <= a < 10^n
	// 10^(1-1/n) <= x < 10, x is a positive integer
	// For 10^(1-1/n) > 9, the set of possibilities for x is empty
	// 1-1/n > log10(9) => 1/n < 1 - log10(9) => n > 1 / (1 - log10(9)) == 21.8 => n < 22 => a < 10^22

	solutions := make(map[float64]struct{})
	for n := 1; n < 22; n++ {
		for x := 1; x < 10; x++ {
			a := math.Pow(float64(x), float64(n))
			if a >= math.Pow10(n-1) && a < math.Pow10(n) {
				solutions[a] = struct{}{}
			}
		}
	}

	result = strconv.FormatInt(int64(len(solutions)), 10)
	return
}
