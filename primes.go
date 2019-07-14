package projecteuler

import (
	"fmt"
	"math"
)

// Primes calculates and returns slice of primes smaller than limit
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

// Powered represents Base^Exp (int) numbers
type Powered struct {
	Base int
	Exp  int
}

// Factorize returns prime factorization of num
func Factorize(num int, primes []int) (factors []Powered, err error) {
	for i := 0; i < len(primes); i++ {
		if num%primes[i] == 0 {
			exp := 0
			for num%primes[i] == 0 {
				num /= primes[i]
				exp++
			}
			factors = append(factors, Powered{Base: primes[i], Exp: exp})

			if num == 1 {
				break
			}
		}
	}

	if num != 1 {
		err = fmt.Errorf("Undivisible residue %d, 'primes' is not adequate", num)
	}

	return
}
