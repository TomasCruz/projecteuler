package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 46; Goldbach's other conjecture

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and
twice a square.

9 = 7 + 2×1^2
15 = 7 + 2×2^2
21 = 3 + 2×3^2
25 = 7 + 2×3^2
27 = 19 + 2×2^2
33 = 31 + 2×1^2

It turns out that the conjecture was false. What is the smallest odd composite that cannot be written as
the sum of a prime and twice a square?
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
		limit = 10000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	roots := projecteuler.ReverseSquares(int(math.Sqrt(float64(limit / 2))))
	primes, primeSet := projecteuler.PrimeSet(limit)

	for i := 9; i < limit; i++ {
		if _, ok := primeSet[i]; ok {
			continue
		}

		var j int
		for j = 1; primes[j] < i; j++ {
			potentialSquare := (i - primes[j]) / 2
			if _, ok := roots[potentialSquare]; ok {
				break
			}
		}

		if primes[j] >= i {
			result = strconv.FormatInt(int64(i), 10)
			break
		}
	}

	return
}
