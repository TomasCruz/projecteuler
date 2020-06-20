package main

import (
	"math"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 37; Truncatable primes

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits
from left to right, and remain prime at each stage: 3797, 797, 97, and 7.
Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/

func main() {
	// I'm cheating here, as I know the solution and limit to truncatable primes (<10^8).
	// So, I'm finding right truncatables, then left truncatable not greater than biggest right truncatable,
	// and sum numbers with both properties
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	limit := 100000
	primes, primeSet := projecteuler.PrimeSet(limit)

	rtp := rightTruncatablePrimes(primes, primeSet)
	biggest := biggestInSet(rtp)
	ltp := leftTruncatablePrimes(primes, primeSet, biggest)

	sum := 0
	for x := range rtp {
		if _, ok := ltp[x]; ok {
			sum += x
		}
	}
	sum -= 17 // deduce one digits results as per requirement

	result = strconv.Itoa(sum)
	return
}

func rightTruncatablePrimes(primes []int, primeSet map[int]struct{}) (rtp map[int]struct{}) {
	// first digit must be prime, hence in {2,3,5,7}
	// other digits must not be even or 5 because of divisibility with 2 and 5, so other digits must be in {1,3,7,9}
	// right truncatables with x+1 digits must, when right truncated, be one of x digit right truncatables

	rtp = make(map[int]struct{})
	rtp[2] = struct{}{}
	rtp[3] = struct{}{}
	rtp[5] = struct{}{}
	rtp[7] = struct{}{}

	lastBatch := []int{2, 3, 5, 7}
	possibleDigits := []int{1, 3, 7, 9}

	for len(lastBatch) > 0 {
		numAdded := 0

		var nextBatch []int
		for i := 0; i < len(lastBatch); i++ {
			base := 10 * lastBatch[i]
			for j := 0; j < len(possibleDigits); j++ {
				curr := base + possibleDigits[j]
				if isPrime(curr, primes, primeSet) {
					rtp[curr] = struct{}{}
					nextBatch = append(nextBatch, curr)
					numAdded++
				}
			}
		}

		lastBatch = make([]int, numAdded)
		copy(lastBatch, nextBatch)
	}

	return
}

func leftTruncatablePrimes(primes []int, primeSet map[int]struct{}, biggest int) (ltp map[int]struct{}) {
	// last digit must be prime, hence in {2,3,5,7}
	// other digits must not be 0
	// left truncatables with x+1 digits must, when left truncated, be one of x digit left truncatables

	ltp = make(map[int]struct{})
	ltp[2] = struct{}{}
	ltp[3] = struct{}{}
	ltp[5] = struct{}{}
	ltp[7] = struct{}{}

	lastPowerTen := 1
	lastBatch := []int{2, 3, 5, 7}
	limitExceded := false

	for !limitExceded {
		numAdded := 0
		lastPowerTen *= 10

		var nextBatch []int
		for i := 0; i < len(lastBatch); i++ {
			for j := 1; j < 10; j++ {
				curr := lastBatch[i] + j*lastPowerTen
				if curr > biggest {
					limitExceded = true
				}

				if isPrime(curr, primes, primeSet) {
					ltp[curr] = struct{}{}
					nextBatch = append(nextBatch, curr)
					numAdded++
				}
			}
		}

		lastBatch = make([]int, numAdded)
		copy(lastBatch, nextBatch)
	}

	return
}

func biggestInSet(numSet map[int]struct{}) int {
	ret := math.MinInt32
	for x := range numSet {
		if x > ret {
			ret = x
		}
	}

	return ret
}

// user is responsible to provide big enough primes slice, so that root(x) isn't bigger than primes in the list
func isPrime(x int, primes []int, primeSet map[int]struct{}) bool {
	if _, ok := primeSet[x]; ok {
		return true
	}

	rt := int(math.Sqrt(float64(x)))

	var i int
	for i = 0; i < len(primes) && primes[i] <= rt; i++ {
		if x%primes[i] == 0 {
			return false
		}
	}

	primeSet[x] = struct{}{}
	return true
}
