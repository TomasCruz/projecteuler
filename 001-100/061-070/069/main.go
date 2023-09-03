package main

import (
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 69; Totient Maximum
Euler's totient function, phi(n) is defined as the number of positive integers not exceeding n which are relatively prime to n.
For example, as 1, 2, 4, 5, 7 and 8, are all less than or equal to nine and relatively prime to nine, phi(9) = 6.

n	Relatively Prime 		phi(n)	n/phi(n)
2 	1						1		2
3 	1,2						2		1.5
4 	1,3						2		2
5 	1,2,3,4					4		1.25
6 	1,5						2		3
7 	1,2,3,4,5,6				6		1.1666...
8 	1,3,5,7					4		2
9 	1,2,4,5,7,8				6		1.5
10 	1,3,7,9					4		2.5

It can be seen that n = 6 produces a maximum for n/phi(n) for n <= 10.
Find the value of n <= 1000000 for which n/phi(n) is a maximum.
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
		limit = 1000001
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	primes := projecteuler.Primes(limit, nil)
	numberDividedByTotient := make([]float64, limit)

	for n := 2; n < limit; n++ {
		factors, err := projecteuler.Factorize(n, primes)
		if err != nil {
			return "", err
		}

		currFactors := []int{}
		for f := range factors {
			currFactors = append(currFactors, f)
		}
		sort.Ints(currFactors)

		totient := projecteuler.Totient(n, currFactors)
		numberDividedByTotient[n] = float64(n) / float64(totient)
	}

	maxN := -1
	maxCoef := float64(-1)
	for n, coef := range numberDividedByTotient {
		if coef > maxCoef {
			maxN = n
			maxCoef = coef
		}
	}

	result = strconv.Itoa(maxN)
	return
}
