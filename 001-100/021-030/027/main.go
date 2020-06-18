package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 27; Quadratic primes

Euler discovered the remarkable quadratic formula:

n^2+n+41

It turns out that the formula will produce 40 primes for the consecutive integer values 0≤n≤39
. However, when n=40,402+40+41=40(40+1)+41 is divisible by 41, and certainly when n=41,412+41+41

is clearly divisible by 41.

The incredible formula n^2−79n+1601
was discovered, which produces 80 primes for the consecutive values 0≤n≤79

. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

    n^2+an+b

, where |a|<1000 and |b|≤1000

where |n|
is the modulus/absolute value of n
e.g. |11|=11 and |−4|=4

Find the product of the coefficients, a and b, for the quadratic expression that
produces the maximum number of primes for consecutive values of n, starting with n=0.
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
		limit = 1000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	// b, 1+a+b, 4+2a+b, 9+3a+b, 16+4a+b, 25+5a+b...
	maxConsecutives, maxA, maxB := 0, 0, 0
	_, primesSet := projecteuler.PrimesSet(1000000)
	_, primesLimitSet := projecteuler.PrimesSet(limit)

	for b := range primesLimitSet {
		currConsecutives, currA := maxNumberOfConsecutives(limit, b, primesSet)
		if currConsecutives > maxConsecutives {
			maxConsecutives = currConsecutives
			maxA = currA
			maxB = b
		}
	}

	resultInt := maxA * maxB
	result = strconv.Itoa(resultInt)
	return
}

func maxNumberOfConsecutives(limit, b int, primesSet map[int]struct{}) (consecutiveCount int, maxA int) {
	consecutiveCount = 1

	for a := 1 - limit; a < limit; a++ {
		curr := numberOfConsecutives(a, limit, b, primesSet)
		if curr > consecutiveCount {
			consecutiveCount = curr
			maxA = a
		}
	}

	return
}

func numberOfConsecutives(a, limit, b int, primesSet map[int]struct{}) (consecutiveCount int) {
	consecutiveCount = 1

	for n := 1; ; n++ {
		expr := n*n + a*n + b
		if _, ok := primesSet[expr]; !ok {
			break
		}

		consecutiveCount++
	}

	return
}
