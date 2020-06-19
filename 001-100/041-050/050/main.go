package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 50; Consecutive prime sum

The prime 41, can be written as the sum of six consecutive primes:
41 = 2 + 3 + 5 + 7 + 11 + 13

This is the longest sum of consecutive primes that adds to a prime below one-hundred. The longest sum of
consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.
Which prime, below one-million, can be written as the sum of the most consecutive primes?
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
		limit = 1000000
	}

	projecteuler.Timed(calc, limit)
}

type primeSum struct {
	prime, terms int
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	primes, primeSet := projecteuler.PrimesSet(limit)
	primeCount := len(primes)
	sumMatrix := make([][]int, primeCount)
	primeSumSet := make(map[primeSum]struct{})

	for i := 0; i < primeCount; i++ {
		sumMatrix[i] = append(sumMatrix[i], 0)
		sumMatrix[i] = append(sumMatrix[i], primes[i])
		lastAdded := sumMatrix[i][1]
		termCount := 1

		for lastAdded < limit {
			if i+termCount >= primeCount {
				break
			}

			sumMatrix[i] = append(sumMatrix[i], lastAdded+primes[i+termCount])
			termCount++
			lastAdded = sumMatrix[i][termCount]

			if _, ok := primeSet[lastAdded]; ok {
				primeSumSet[primeSum{prime: lastAdded, terms: termCount}] = struct{}{}
			}
		}
	}

	biggestSumPrime := 0
	maxTermCount := 0
	for pr := range primeSumSet {
		if pr.terms > maxTermCount {
			maxTermCount = pr.terms
			biggestSumPrime = pr.prime
			continue
		}

		if pr.terms == maxTermCount && pr.prime > biggestSumPrime {
			biggestSumPrime = pr.prime
		}
	}

	result = strconv.Itoa(biggestSumPrime)
	return
}
