package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 41; Pandigital prime

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once.
For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	// m-digit pandigital prime has sum of digits 1+2+...+m = 1/2*m*(m+1)
	// for m being 9 or 8 sum is divisible by 3 and number is not prime.
	// So, largest n-digit pandigital prime has 7 digits (or less, but I already know it's 7 :)

	var res int
	projecteuler.Permutations(7, isLargest, &res)
	result = strconv.Itoa(res)
	return
}

func isLargest(args ...interface{}) bool {
	resPtr := args[0].(*int)
	currPerm := args[1].([]byte)

	res := 0
	digits := len(currPerm)
	for i := 0; i < digits; i++ {
		res *= 10
		res += digits - int(currPerm[i])
	}

	if !isPrime(res) {
		return false
	}

	*resPtr = res
	return true
}

func isPrime(x int) bool {
	if x%2 == 0 {
		return false
	}
	div := 3
	limit := x / 3

	for div < limit {
		if x%div == 0 {
			return false
		}
		div += 2
	}

	return true
}
