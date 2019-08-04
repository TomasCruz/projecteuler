package projecteuler

import (
	"fmt"
)

// Factorize returns prime factorization of num
func Factorize(num int, primes []int) (factors map[int]int, err error) {
	factors = make(map[int]int)

	for i := 0; i < len(primes); i++ {
		if num%primes[i] == 0 {
			exp := 0
			for num%primes[i] == 0 {
				num /= primes[i]
				exp++
			}
			factors[primes[i]] = exp

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

// MultiplyFactors multiplies factors
func MultiplyFactors(primeFactors map[int]int) int64 {
	result := int64(1)

	for k, v := range primeFactors {
		for j := 0; j < v; j++ {
			result *= int64(k)
		}
	}

	return result
}

// CompareFactors comapares factors of two numbers returning true if equal
func CompareFactors(f1, f2 map[int]int) bool {
	if len(f1) != len(f2) {
		return false
	}

	for k, v := range f1 {
		var v2 int
		var ok bool

		if v2, ok = f2[k]; !ok {
			return false
		}

		if v != v2 {
			return false
		}
	}

	return true
}
