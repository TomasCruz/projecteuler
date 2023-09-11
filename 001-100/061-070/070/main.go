package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 70;
Euler's totient function, phi(n) [sometimes called the phi function], is used to determine the number of positive numbers
less than or equal to n which are relatively prime to n. For example, as 1, 2, 4, 5, 7 and 8, are all less than or equal
to nine and relatively prime to nine, phi(9) = 6.
The number 1 is considered to be relatively prime to every positive number, so phi(1) = 1.

Interestingly, phi(87109) = 79180, and it can be seen that 87109 is a permutation of 79180.

Find the value of n, 1 < n < 10^7, for which phi(n) is a permutation of n and the ratio n/phi(n) produces a minimum.
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
		limit = 10000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	primes := projecteuler.Primes(limit, nil)

	minPrimeIndex := 20
	for ; primes[minPrimeIndex] < 1000; minPrimeIndex++ {
	}

	lenPrimes := len(primes)
	minX := 0
	minRatio := math.MaxFloat64

	for i := minPrimeIndex; i < lenPrimes; i++ {
		for j := i + 1; j < lenPrimes; j++ {
			x := primes[i] * primes[j]
			if x >= limit {
				break
			}

			tot := totient(i, j, primes)
			if permutation(x, tot) {
				ratio := float64(x) / float64(tot)
				if ratio < minRatio {
					minX = x
					minRatio = ratio
				}
			}
		}
	}

	result = strconv.Itoa(minX)
	return
}

func totient(i, j int, primes []int) int {
	return (primes[i] - 1) * (primes[j] - 1)
}

func permutation(a, b int) bool {
	aDigits := []int{}
	for a > 0 {
		aDigits = append(aDigits, a%10)
		a /= 10
	}
	aLen := len(aDigits)

	bDigits := []int{}
	for b > 0 {
		bDigits = append(bDigits, b%10)
		b /= 10
	}
	bLen := len(bDigits)

	if aLen != bLen {
		return false
	}

	sort.Ints(aDigits)
	sort.Ints(bDigits)

	for i := 0; i < aLen; i++ {
		if aDigits[i] != bDigits[i] {
			return false
		}
	}

	return true
}
