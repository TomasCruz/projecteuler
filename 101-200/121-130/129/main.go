package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 129; Repunit Divisibility
A number consisting entirely of ones is called a repunit. We shall define R(k) to be a repunit of length k; for example, R(6) = 111111.

Given that n is a positive integer and gcd(n, 10) = 1, it can be shown that there always exists a value, k, for which R(k) is divisible by n,
and let A(n) be the least such value of k; for example, A(7) = 6 and A(41) = 5.

The least value of n for which A(n) first exceeds ten is 17.
Find the least value of n for which A(n) first exceeds one-million.
*/

func main() {
	var limit int

	if len(os.Args) > 1 {
		limit64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limit = int(limit64)
	} else {
		limit = 6
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	power := 1
	for range limit {
		power *= 10
	}

	maxOrder := int(float64(limit) / math.Log10(3.0))
	primes := projecteuler.Primes(2*power, nil)
	multiplicativeOrderPrimes := make([][]map[int]int, len(primes))
	for i := range primes {
		multiplicativeOrderPrimes[i] = make([]map[int]int, maxOrder+3)
	}

	n := 0
	base := power
	mods := []int{1, 3, 7, 9}

outter:
	for {
		for _, j := range mods {
			n = base + j
			var moMap map[int]int
			moMap, multiplicativeOrderPrimes = multiplicativeOrder10Modulo(9*n, primes, multiplicativeOrderPrimes)

			prod := 1
			for k, v := range moMap {
				for range v {
					prod *= primes[k]
				}
			}

			if prod > power {
				break outter
			}
		}

		base += 10
	}

	result = strconv.Itoa(n)
	return
}

func multiplicativeOrder10Modulo(n int, primes []int, multiplicativeOrderPrimes [][]map[int]int) (map[int]int, [][]map[int]int) {
	nFactors := projecteuler.FactorizeIndex(n, primes)
	factorMaps := make([]map[int]int, 0, len(nFactors))
	for k, v := range nFactors {
		if multiplicativeOrderPrimes[k][v] == nil {
			multiplicativeOrderPrimes[k][v] = calcMulOrder10ModuloForPrimePower(primes[k], v, primes)
		}

		factorMaps = append(factorMaps, multiplicativeOrderPrimes[k][v])
	}

	lcmFactorMap := lcm(factorMaps)
	return lcmFactorMap, multiplicativeOrderPrimes
}

func calcMulOrder10ModuloForPrimePower(p, power int, primes []int) map[int]int {
	pp := 1
	for range power {
		pp *= p
	}

	pMod := 10 % pp
	pow10 := pMod
	i := 1
	for pow10 != 1 {
		pow10 = (pow10 * pMod) % pp
		i++
	}

	return projecteuler.FactorizeIndex(i, primes)
}

func lcm(numbers []map[int]int) map[int]int {
	factorMap := map[int]int{}

	for i := range numbers {
		for k, v := range numbers[i] {
			currV, present := factorMap[k]
			if present {
				if v > currV {
					factorMap[k] = v
				}
			} else {
				factorMap[k] = v
			}
		}
	}

	return factorMap
}

/*
	Input from users DJohn and ilarrosa:

	User DJohn

	A(n) <= n:

	Consider the first n+1 repunits: R(1) = 1, R(2) = 11, ..., R(n+1) = 111...111.

	When divided by n, they give n+1 remainders all between 0 and n-1 (at most n different remainders).  By the Dirichlet's principle (pigeonhole),
	at least one pair of these remainders must be equal.

	Let r and s be the repunits that give a pair of equal remainders, with r > s.  The remainder from dividing r-s by n is zero; that is, r-s is divisible by n.

	r-s is of the form 111...111000...000.  n is coprime with 10, so we can divide r-s by a power of ten (say m) to remove the trailing zeroes.
	That leaves us with a repunit.  This repunit must be less than r.

	We only considered the first n+1 repunits, so r must be at most R(n+1).  Since (r-s)/10^m < r, it can be at most R(n).
	So any n coprime with 10 is a divisor of at least one of the first n repunits.

	User ilarrosa

	R(k)= (10^k - 1)/9
	10^k - 1 = 0 (mod n)
	It is true of course for k = phi(n), and if it is true for another k < phi(n), it must be k | phi(n).
	If n = 0 (mod 3), then it must be 10^k - 1 = 0 (mod 9n) and k | phi(9n)
*/
