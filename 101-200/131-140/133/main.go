package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 133; Repunit Nonfactors
A number consisting entirely of ones is called a repunit. We shall define R(k) to be a repunit of length k; for example, R(6) = 111111.

Let us consider repunits of the form R(10^n).

Although R(10), R(100), or R(1000) are not divisible by 17, R(10000) is divisible by 17.
Yet there is no value of n for which R(10^n) will divide by 19. In fact, it is remarkable that 11, 17, 41, and 73
are the only four primes below one-hundred that can be a factor of R(10^n).

Find the sum of all the primes below one-hundred thousand that will never be a factor of R(10^n).
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
		limit = 100000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(limit, nil)

	multiplicativeOrderPrimes := make([][]map[int]int, len(primes))
	for i := range primes {
		multiplicativeOrderPrimes[i] = make([]map[int]int, 21)
	}

	sum := int64(10) // 2, 3, 5
	for i := 3; i < len(primes); i++ {
		var mo map[int]int
		mo, multiplicativeOrderPrimes = multiplicativeOrder10Modulo(primes[i], primes, multiplicativeOrderPrimes)
		if !divisible(mo) {
			sum += int64(primes[i])
		}
	}

	result = strconv.FormatInt(sum, 10)
	return
}

func divisible(mo map[int]int) bool {
	for k := range mo {
		if k != 0 && k != 2 {
			return false
		}
	}

	return true
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
