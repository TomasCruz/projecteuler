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
	primes, primeSet := projecteuler.PrimeSet(limit)
	primeCount := len(primes)
	sumMatrix := make([][]int, primeCount)

	// dynamically build sumMatrix
	sumMatrix[primeCount-1] = append(sumMatrix[primeCount-1], primes[primeCount-1])
	for i := primeCount - 1; i > 0; i-- {
		sumMatrix[i-1] = append(sumMatrix[i-1], primes[i-1])

		for j := 0; j < len(sumMatrix[i]); j++ {
			sum := sumMatrix[i][j] + primes[i-1]
			if sum >= limit {
				break
			}

			sumMatrix[i-1] = append(sumMatrix[i-1], sum)
		}
	}

	// find the greatest sum, having maximum number of terms
	biggestSumPrime := 0
	maxTermCount := 0
	for i := 0; i < primeCount; i++ {
		for j := len(sumMatrix[i]); j > 0; j-- {
			currSum := sumMatrix[i][j-1]
			if _, ok := primeSet[currSum]; ok {
				if j > maxTermCount {
					maxTermCount = j
					biggestSumPrime = currSum
				}

				if j == maxTermCount && currSum > biggestSumPrime {
					biggestSumPrime = currSum
				}

				// this prime sum is the greatest sum to start with prime[i], so move on to next prime
				break
			}
		}
	}

	result = strconv.Itoa(biggestSumPrime)
	return
}
