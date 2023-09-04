package projecteuler

func Totient(n int, primes []int) int {
	result := n

	// Consider all prime factors of n and for every prime factor p, and subtract their multiples from result
	for i := 0; primes[i]*primes[i] <= n; i++ {
		if n%primes[i] == 0 {
			for n%primes[i] == 0 {
				n /= primes[i]
			}
			result -= result / primes[i]
		}
	}

	if n > 1 {
		result -= result / n
	}

	return result
}

func NumDividedByTotient(n int, primes []int) float64 {
	return float64(n) / float64(Totient(n, primes))
}
