package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 132; Large Repunit Factors
A number consisting entirely of ones is called a repunit. We shall define R(k) to be a repunit of length k.

For example, R(10) = 1111111111 = 11 * 41 * 271 * 9091, and the sum of these prime factors is 9414.

Find the sum of the first forty prime factors of R(10^9).
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
		limit = 40
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	factors := map[int]int{0: 9, 2: 9}
	primes := projecteuler.Primes(int(2e5), nil)

	multiplicativeOrderPrimes := make([][]map[int]int, len(primes))
	for i := range primes {
		multiplicativeOrderPrimes[i] = make([]map[int]int, 21)
	}

	count := 0
	sum := 0
	for i := 3; count < limit; i++ {
		var mo map[int]int
		mo, multiplicativeOrderPrimes = multiplicativeOrder10Modulo(primes[i], primes, multiplicativeOrderPrimes)
		if divisible(mo, factors) {
			count++
			sum += primes[i]
		}
	}

	result = strconv.Itoa(sum)
	return
}

func divisible(mo map[int]int, factors map[int]int) bool {
	for k, v := range mo {
		v1, present := factors[k]
		if !present || v1 < v {
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
