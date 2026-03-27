package projecteuler

import (
	"fmt"
	"slices"
)

// Factorize returns prime factorization of num, key of the map being prime factor.
// Returns error if primes doesn't contain all prime factors of num
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

// FactorizeIndex returns prime factorization of num, key of the map being index of the prime factor.
// Returns error if primes doesn't contain all prime factors of num
func FactorizeIndex(num int, primes []int) map[int]int {
	factors := make(map[int]int)

	for i := 0; i < len(primes); i++ {
		if num%primes[i] == 0 {
			exp := 0
			for num%primes[i] == 0 {
				num /= primes[i]
				exp++
			}
			factors[i] = exp

			if num == 1 {
				break
			}
		}
	}

	return factors
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

// FindDivisors takes factors map and returns slice of divisors
func FindDivisors(factors map[int]int) (divisors []int) {
	var firstK, firstV int

	otherFactors := make(map[int]int)
	for k, v := range factors {
		if firstK == 0 {
			firstK = k
			firstV = v
		} else {
			otherFactors[k] = v
		}
	}

	powers := make([]int, firstV+1)
	powers[0] = 1
	for i := 0; i < firstV; i++ {
		powers[i+1] = powers[i] * firstK
	}

	if len(otherFactors) == 0 {
		return powers
	}

	otherDivisors := FindDivisors(otherFactors)
	otherDivisorCount := len(otherDivisors)
	divisors = make([]int, (firstV+1)*otherDivisorCount)

	for i := 0; i <= firstV; i++ {
		for j := 0; j < otherDivisorCount; j++ {
			divisors[i*otherDivisorCount+j] = powers[i] * otherDivisors[j]
		}
	}

	return
}

type factorization struct {
	unfactored int
	factors    map[int]int
}

// FactorizeAll returns prime factorization of all the integers less than limit, all the primes less than limit, and all the prime powers less than limit
func FactorizeAll(limit int) ([]map[int]int, []int, [][]int64) {
	factorizations := make([]factorization, limit)

	for i := 2; i < limit; i++ {
		factorizations[i] = factorization{
			unfactored: i,
			factors:    map[int]int{},
		}
	}

	for num := 2; num < limit; num++ {
		if factorizations[num].unfactored == 1 {
			continue
		}

		for multiple := num; multiple < limit; multiple += num {
			exp := 0
			for factorizations[multiple].unfactored%num == 0 {
				factorizations[multiple].unfactored /= num
				exp++
			}
			factorizations[multiple].factors[num] = exp
		}
	}

	factors := make([]map[int]int, limit)

	var primes []int
	var primePowers [][]int64
	if limit <= 10000 {
		l := limit / 2
		primes = make([]int, 0, l)
		primePowers = make([][]int64, 0, l)
	} else {
		l := limit / 10
		primes = make([]int, 0, l)
		primePowers = make([][]int64, 0, l)
	}

	for num := 2; num < limit; num++ {
		factors[num] = make(map[int]int, len(factorizations[num].factors))

		if len(factorizations[num].factors) == 1 {
			for k, v := range factorizations[num].factors {
				if v == 1 {
					primes = append(primes, num)
					primePowers = append(primePowers, []int64{})
					l := len(primePowers) - 1
					primePowers[l] = append(primePowers[l], 1)
					primePowers[l] = append(primePowers[l], int64(num))
					factors[num][l] = 1
				} else {
					kIndex, _ := slices.BinarySearch(primes, k)
					primePowers[kIndex] = append(primePowers[kIndex], int64(num))
					factors[num][kIndex] = v
				}
			}
		} else {
			for k, v := range factorizations[num].factors {
				kIndex, _ := slices.BinarySearch(primes, k)
				factors[num][kIndex] = v
			}
		}
	}

	return factors, primes, primePowers
}
