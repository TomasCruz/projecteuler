package projecteuler

import (
	"math"
	"unsafe"
)

type Int32Plus interface {
	~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64
}

// PrimesDivisibility calculates and returns slice of primes smaller than limit, or until f returns true
func PrimesDivisibility[T Int32Plus](limit T, f func(...any) bool, args ...any) (primes []T) {
	primes = append(primes, 2)
	primes = append(primes, 3)

	var i int
	for num := T(5); num < limit; num += 2 {
		numRootLimit := T(math.Sqrt(float64(num)))

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
func PrimeSet[T Int32Plus](limit T) (primes []T, primeSet map[T]struct{}) {
	primes = Primes(limit, nil)

	primeSet = make(map[T]struct{})
	for _, x := range primes {
		primeSet[x] = struct{}{}
	}

	return
}

// IsPrime returns true iff x is prime
func IsPrime[T Int32Plus](x T) bool {
	if x%2 == 0 {
		return false
	}

	root := T(math.Sqrt(float64(x)))
	for i := T(3); i <= root; i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}

// Primes uses Sieve of Eratosthenes to calculates and returns slice of primes smaller than limit, or until f returns true
func Primes[T Int32Plus](limit T, f func(...any) bool, args ...any) (primes []T) {
	bitsize := int(8 * unsafe.Sizeof(limit))
	bs := NewBitset(limit+1, bitsize)

	limitRoot := int(math.Sqrt(float64(limit)))
	for i := 2; i <= limitRoot; i++ {
		if bs.Get(i) {
			continue
		}

		for j := 2 * i; j < int(limit); j += i {
			bs.Set(j, true)
		}
	}

	var ret []T
	if limit <= T(10000) {
		ret = make([]T, 0, limit/2)
	} else {
		ret = make([]T, 0, limit/10)
	}

	ret = append(ret, T(2))

	for i := 3; i <= int(limit); i += 2 {
		if bs.Get(i) {
			continue
		}

		ret = append(ret, T(i))
		if f != nil {
			args = append(args, ret)
			if f(args...) {
				break
			}
			args = args[:len(args)-1]
		}
	}

	return ret
}

// PrimePowers takes primes less than limit and returns prime powers less than limit, for powers less than powerLimit
func PrimePowers[T Int32Plus](primes []T, limit, powerLimit int) [][]int64 {
	primePowers := make([][]int64, len(primes))

	for i, currPrime := range primes {
		primePowers[i] = make([]int64, 0, powerLimit)
		primePowers[i] = append(primePowers[i], 1)

		for j := 1; j < powerLimit; j++ {
			prev := primePowers[i][j-1]
			next := prev * int64(currPrime)
			if next > int64(limit) {
				break
			}

			primePowers[i] = append(primePowers[i], next)
		}
	}

	return primePowers
}
