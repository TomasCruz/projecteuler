package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 53; Combinatoric selections

There are exactly ten ways of selecting three from five, 12345:
123, 124, 125, 134, 135, 145, 234, 235, 245, and 345
In combinatorics, we use the notation, C(5, 3)=10
In general, C(n, r)=n!/[r!(n−r)!], where r <= n, n!=n×(n−1)×...×3×2×1, and 0!=1

It is not until n=23, that a value exceeds one-million: C(23, 10)=1144066
How many, not necessarily distinct, values of C(n, r) for 1<=n<=100, are greater than one-million?
*/

func main() {
	var limit, threshold int

	if len(os.Args) > 2 {
		limit64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limit = int(limit64)

		threshold64, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		threshold = int(threshold64)
	} else {
		limit = 100
		threshold = 1000000
	}

	projecteuler.Timed(calc, limit, threshold)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	threshold := args[1].(int)

	primes := projecteuler.Primes(limit+1, nil)
	result64 := int64(0)
	for i := 1; i <= limit; i++ {
		for j := 2; j < i; j++ {
			var primeFactors map[int]int
			if primeFactors, err = projecteuler.Binomial(i, j, primes); err != nil {
				return
			}

			if overThreshold(threshold, primeFactors) {
				result64++
			}
		}
	}
	result = strconv.FormatInt(result64, 10)

	return
}

func overThreshold(threshold int, primeFactors map[int]int) bool {
	num := int64(1)
	for prime, occ := range primeFactors {
		for i := 0; i < occ; i++ {
			num *= int64(prime)
			if num > int64(threshold) {
				return true
			}
		}
	}

	return false
}
