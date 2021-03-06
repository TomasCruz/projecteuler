package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 21; Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
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
		limit = 10000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(limit, nil)
	properDivisors := make([]int, limit)

	for i := 4; i < limit; i++ {
		var currFactors map[int]int
		if currFactors, err = projecteuler.Factorize(i, primes); err != nil {
			return
		}

		properDivisors[i] = divisorSum(currFactors) - i
	}

	amicableSum := 0
	for i := 4; i < limit-1; i++ {
		pds := properDivisors[i]
		if pds > i && pds < limit && properDivisors[pds] == i {
			amicableSum += i + pds
		}
	}

	result = strconv.Itoa(amicableSum)
	return
}

func divisorSum(factors map[int]int) (sum int) {
	divisors := projecteuler.FindDivisors(factors)

	for i := 0; i < len(divisors); i++ {
		sum += divisors[i]
	}

	return
}
