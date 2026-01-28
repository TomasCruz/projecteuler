package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 110; Diophantine Reciprocals II
In the following equation x, y, and n are positive integers.

1/x + 1/y = 1/n

It can be verified that when n = 1260 there are 113 distinct solutions and this is the least value of n for which the total number of distinct solutions
exceeds one hundred.

What is the least value of n for which the number of distinct solutions exceeds four million?

NOTE: This problem is a much more difficult version of Problem 108 and as it is well beyond the limitations of a brute force approach it requires
a clever implementation.
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
		limit = 4000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	n := int64(1)
	fc := newFC(limit)
	maxPrimeCount := int(math.Ceil(math.Log10(float64(limit)) / math.Log10(3.0)))
	maxFactorSubsetMask := make([]int, maxPrimeCount)
	dontExceedFactors := make([]int, maxPrimeCount)

	for i := range maxPrimeCount {
		n *= fc.primePowers[i][1]
		maxFactorSubsetMask[i] = 1
		dontExceedFactors[i] = len(fc.primePowers[i]) - 1
	}

	fts := fc.calcSubsetFactorTypes(n, maxFactorSubsetMask, dontExceedFactors)

	var currFT factorType
	for _, currFT = range fts {
		divisorCount := fc.findDivisorsLesserThanRootCount(currFT)
		if divisorCount > limit {
			break
		}
	}

	result = strconv.FormatInt(currFT.n, 10)
	return
}

func buidPrimePowers(primes []int) [][]int64 {
	primePowers := make([][]int64, len(primes))

	for i, currPrime := range primes {
		primePowers[i] = make([]int64, 0, 50)
		primePowers[i] = append(primePowers[i], 1)

		for len(primePowers[i]) < 50 {
			prev := primePowers[i][len(primePowers[i])-1]
			next := prev * int64(currPrime)
			if next < prev { // overflow
				break
			}
			primePowers[i] = append(primePowers[i], next)
		}
	}

	return primePowers
}

/*
	x > n
	for x <= y, 1/x >= 1/y => 1/x + 1/y <= 2/x, 1/n <= 2/x, x <= 2n

	x in [n+1,...,2n]
	x-n in [1,...,n]

	1/n = 1/x + 1/y / *nxy
	nx + ny = yx
	nx = y*(x - n)

	y = nx / (x-n) = n*(x-n+n) / (x-n) = n + n^2 / (x-n) => y-n = n^2 / (x-n)

	Considering (x-n)*(y-n) = n^2, x-n and y-n have to be "complementary" divisors of n^2.
	So, A = x-n and B = y-n, and A*B = n^2.
	One solution for (a, b) is always (n, n).
	Let c = number of As, i.e. number of divisors of d^2 which are <= n.
	c-1 is the number of divisors of d^2 which are < n, so number of Bs > n is also c-1.
	Therefore number of divisors of d^2 which appear in solutions (a, b) is 2*(c-1)+1 = 2*c-1.
	2*c-1 <= d(n^2) and c > 4*10^6 => d(n^2) > 8*10^6 - 1, i.e. d(n^2) >= 2 * limit

	P(ai + 1), 3^k > 4*10^6
	k > (6 + log4)/log3 = 13.8, k >= 14
*/
