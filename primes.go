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

type Bitset[T Int32Plus] struct {
	Slice   []T
	Bitsize int
}

// NewBitset construct new Bitset, all the bits set to 0 (false)
func NewBitset[T Int32Plus](n T, bitsize int) Bitset[T] {
	return Bitset[T]{
		Slice:   make([]T, (int(n)+bitsize-1)/bitsize),
		Bitsize: bitsize,
	}
}

// Get returns bool value on index
func (b Bitset[T]) Get(index int) bool {
	pos := index / b.Bitsize
	j := index % b.Bitsize
	return (b.Slice[pos] & (T(1) << j)) != 0
}

// All returns set of values in Bitset
func (b Bitset[T]) All() map[int]struct{} {
	m := map[int]struct{}{}

	nPos := 0
	for pos := 0; pos < len(b.Slice); pos++ {
		if b.Slice[pos] == 0 {
			continue
		}

		bit := T(1)
		for j := 0; j < b.Bitsize; j++ {
			if b.Slice[pos]&bit != 0 {
				m[nPos+j] = struct{}{}
			}
			bit <<= 1
		}

		nPos += b.Bitsize
	}

	return m
}

// Set sets value on index
func (b Bitset[T]) Set(index int, value bool) {
	pos := index / b.Bitsize
	j := index % b.Bitsize

	if value {
		b.Slice[pos] |= T(1) << j
	} else {
		// m := uint64(math.MaxUint64)
		// m -= uint64(1) << j
		// b.Slice[pos] &= m
		k := j + 1
		m := ((T(1) << k) - 1) & b.Slice[pos]
		b.Slice[pos] = ((b.Slice[pos] >> k) << k) + m
	}
}

// Primes uses Sieve of Eratosthenes to calculates and returns slice of primes smaller than limit, or until f returns true
func Primes[T Int32Plus](limit T, f func(...any) bool, args ...any) (primes []T) {
	bitsize := int(8 * unsafe.Sizeof(limit))
	bs := NewBitset(limit+1, bitsize)

	for i := 2; i < int(limit); i++ {
		if bs.Get(i) {
			continue
		}

		for j := 2 * i; j < int(limit); j += i {
			bs.Set(j, true)
		}
	}

	var ret []T
	if limit <= T(100000) {
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
