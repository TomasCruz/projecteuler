package projecteuler

import (
	"math"
)

// Primes calculates and returns slice of primes smaller than limit, or until f returns true
func Primes(limit int, f func(...interface{}) bool, args ...interface{}) (primes []int) {
	primes = append(primes, 2)
	primes = append(primes, 3)

	var i, num, numRootLimit int
	for num = 5; num < limit; num += 2 {
		numRootLimit = int(math.Sqrt(float64(num)))

		for i = 0; primes[i] <= numRootLimit; i++ {
			if num%primes[i] == 0 {
				break
			}
		}

		if primes[i] > numRootLimit {
			primes = append(primes, num)
			if f != nil {
				args = append(args, num)
				if f(args...) {
					break
				}
				args = args[:len(args)-1]
			}
		}
	}

	return
}

// PrimeSet calculates and returns slice of primes less than limit and puts them in primeSet
func PrimeSet(limit int) (primes []int, primeSet map[int]struct{}) {
	primes = Primes(limit, nil)

	primeSet = make(map[int]struct{})
	for _, x := range primes {
		primeSet[x] = struct{}{}
	}

	return
}
