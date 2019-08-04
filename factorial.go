package projecteuler

import "fmt"

// Factorial returns factorial of n as map of its prime factors
func Factorial(n int, primes []int) (factors map[int]int, err error) {
	if n < 0 {
		err = fmt.Errorf("negative argument to Factorial")
		return
	}

	factors = make(map[int]int)
	if n == 0 || n == 1 {
		return
	}

	for i := 2; i <= n; i++ {
		var currFactors map[int]int
		if currFactors, err = Factorize(i, primes); err != nil {
			return
		}

		for k, v := range currFactors {
			if _, ok := factors[k]; ok {
				factors[k] += v
			} else {
				factors[k] = v
			}
		}
	}

	return
}

// Binomial return binomial coefficient of n choose k
func Binomial(n, k int, primes []int) (result map[int]int, err error) {
	var factN, factK, factNK map[int]int

	if factN, err = Factorial(n, primes); err != nil {
		return
	}

	if factK, err = Factorial(k, primes); err != nil {
		return
	}

	if factNK, err = Factorial(n-k, primes); err != nil {
		return
	}

	result = make(map[int]int)
	for k, v := range factN {
		if vK, ok := factK[k]; ok {
			v -= vK
		}

		if vNK, ok := factNK[k]; ok {
			v -= vNK
		}

		if v != 0 {
			result[k] = v
		}
	}

	return
}
