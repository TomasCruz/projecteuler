package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 47; Distinct primes factors

The first two consecutive numbers to have two distinct prime factors are:
14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:
644 = 2^2 × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors each.
What is the first of these numbers?
*/

func main() {
	var limit int

	if len(os.Args) > 1 {
		limit64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limit = int(limit64)
	} else {
		limit = 150000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	primes := projecteuler.Primes(limit, nil)

	i := 1
	numPrimeFactors := 4
	consecutiveFound := 0

	for i = 1; i < limit; i++ {
		if hasExpectedDistinctPrimeFactorsCount(i, numPrimeFactors, primes) {
			consecutiveFound++
			if consecutiveFound == numPrimeFactors {
				result = strconv.FormatInt(int64(i-numPrimeFactors+1), 10)
				break
			}
		} else {
			consecutiveFound = 0
		}
	}

	if i == limit {
		result = "not found (404)"
	}

	return
}

func hasExpectedDistinctPrimeFactorsCount(x, numPrimeFactors int, primes []int) bool {
	factors, err := projecteuler.Factorize(x, primes)
	if err != nil {
		panic(x)
	}

	return len(factors) == numPrimeFactors
}
