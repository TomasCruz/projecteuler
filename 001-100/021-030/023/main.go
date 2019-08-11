package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 23; Non-abundant sums

A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant
if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two
abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the
sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known
that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
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
		limit = 28124
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	// for limit == 25, sum = 1..23 = 23*12 = 230+46 = 276
	primes := projecteuler.Primes(limit, nil)
	properDivisors := make([]int, limit)

	for i := 4; i < limit; i++ {
		var currFactors map[int]int
		if currFactors, err = projecteuler.Factorize(i, primes); err != nil {
			return
		}

		properDivisors[i] = divisorSum(currFactors) - i
	}

	abundant := make([]int, 0)
	abundantSet := make(map[int]struct{})
	for i := 4; i < limit; i++ {
		if properDivisors[i] > i {
			abundant = append(abundant, i)
			abundantSet[i] = struct{}{}
		}
	}

	nonAbundantSum := make([]int, 0)
	for i := 1; i < limit; i++ {
		if isNonAbundantSum(i, abundant, abundantSet) {
			nonAbundantSum = append(nonAbundantSum, i)
		}
	}

	totalSum := 0
	for _, n := range nonAbundantSum {
		totalSum += n
	}

	result = strconv.Itoa(totalSum)
	return
}

func divisorSum(factors map[int]int) (sum int) {
	divisors := projecteuler.FindDivisors(factors)

	for i := 0; i < len(divisors); i++ {
		sum += divisors[i]
	}

	return
}

func isNonAbundantSum(x int, abundant []int, abundantSet map[int]struct{}) bool {
	for i := 0; i < len(abundant) && abundant[i] < x; i++ {
		rest := x - abundant[i]
		if rest < abundant[i] {
			break
		}

		if _, ok := abundantSet[rest]; ok {
			return false
		}
	}

	return true
}
