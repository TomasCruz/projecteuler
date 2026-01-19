package main

import (
	"math"
	"sort"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 104; Pandigital Fibonacci Ends
The Fibonacci sequence is defined by the recurrence relation:

F_n = F_{n - 1} + F_{n - 2}, where F_1 = 1 and F_2 = 1.

It turns out that F_{541}, which contains 113 digits, is the first Fibonacci number for which the last nine digits are 1-9 pandigital
(contain all the digits 1 to 9, but not necessarily in order). And F_{2749}, which contains 575 digits, is the first Fibonacci number for which the first nine digits
are 1-9 pandigital.

Given that F_k is the first Fibonacci number for which the first nine digits AND the last nine digits are 1-9 pandigital, find k.
*/

func main() {
	projecteuler.Timed(calc)
}

var logPhi = math.Log10(math.Phi)
var logSqrt5 = math.Log10(math.Sqrt(float64(5)))

func calc(args ...interface{}) (result string, err error) {
	f0 := 0
	f1 := 1
	for range 40 {
		f0, f1 = f1, f0+f1
	}

	k := 41
	for {
		k++
		f0, f1 = f1, (f0+f1)%1e9
		if isPandigital(f1) && isPandigital(firstNineDigitsNthFib(k)) {
			break
		}
	}

	result = strconv.Itoa(k)
	return
}

func firstNineDigitsNthFib(n int) int {
	// Fn == math.Round(Phi^n/sqrt(5)) => log(Fn) == n * log(Phi) - log(sqrt(5))
	x := float64(n)*logPhi - logSqrt5
	digs := math.Floor(math.Pow(10.0, x-math.Floor(x)+8.0))

	return int(digs)
}

func isPandigital(x int) bool {
	digits := make([]int, 9)

	for i := 0; i < 9; i++ {
		digits[i] = x % 10
		x /= 10
	}

	sort.Ints(digits)

	var i int
	for i = 0; i < 9 && digits[i] == i+1; i++ {
	}

	return i == 9
}
