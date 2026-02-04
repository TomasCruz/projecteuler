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

// IsPrime returns true iff x is prime
func IsPrime(x int64) bool {
	if x%2 == 0 {
		return false
	}

	root := int64(math.Sqrt(float64(x)))
	for i := int64(3); i <= root; i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}

type Bitset []uint64

// NewBitset construct new Bitset, all the bits set to 0 (false)
func NewBitset(n uint64) Bitset {
	return make(Bitset, (n+63)/64)
}

func (b Bitset) Get(index uint64) bool {
	pos := index / 64
	j := index % 64
	return (b[pos] & (uint64(1) << j)) != 0
}

func (b Bitset) Set(index uint64, value bool) {
	pos := index / 64
	j := index % 64

	if value {
		b[pos] |= uint64(1) << j
	} else {
		m := uint64(math.MaxUint64)
		m -= uint64(1) << j
		b[pos] &= m
	}
}

// PrimesEratosthenes calculates and returns slice of primes smaller than limit, or until f returns true
func PrimesEratosthenes(limit uint64, f func(...interface{}) bool, args ...interface{}) (primes []uint64) {
	bs := NewBitset(limit + 1)

	for i := uint64(2); i < limit; i++ {
		if bs.Get(i) {
			continue
		}

		for j := uint64(2 * i); j < limit; j += i {
			bs.Set(j, true)
		}
	}

	var ret []uint64
	if limit <= uint64(100000) {
		ret = make([]uint64, 0, limit/2)
	} else {
		ret = make([]uint64, 0, limit/10)
	}

	ret = append(ret, uint64(2))

	for i := uint64(3); i <= limit; i += 2 {
		if bs.Get(i) {
			continue
		}

		ret = append(ret, i)
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
