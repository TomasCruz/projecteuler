package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 88; Product-sum Numbers
A natural number, N, that can be written as the sum and product of a given set of at least two natural numbers, {a1, a2, ... , ak}
is called a product-sum number: N = a1 + a2 + ... + ak = a1 * a2 * ... * ak. For example, 6 = 1 + 2 + 3 = 1 * 2 * 3.

For a given set of size, k, we shall call the smallest N with this property a minimal product-sum number. The minimal product-sum numbers
for sets of size, k = 2, 3, 4, 5, and 6 are as follows.
k = 2: 4 = 2 * 2 = 2 + 2
k = 3: 6 = 1 * 2 * 3 = 1 + 2 + 3
k = 4: 8 = 1 * 1 * 2 * 4 = 1 + 1 + 2 + 4
k = 5: 8 = 1 * 1 * 2 * 2 * 2 = 1 + 1 + 2 + 2 + 2
k = 6: 12 = 1 * 1 * 1 * 1 * 2 * 6 = 1 + 1 + 1 + 1 + 2 + 6

Hence for 2 <= k <= 6, the sum of all the minimal product-sum numbers is 4 + 6 + 8 + 12 = 30; note that 8 is only counted once in the sum.
In fact, as the complete set of minimal product-sum numbers for 2 <= k <= 12 is {4, 6, 8, 12, 15, 16}, the sum is 61.
What is the sum of all the minimal product-sum numbers for 2 <= k <= 12000?
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
		limit = 12000
	}

	projecteuler.Timed(calc, limit)
}

type sumTerm struct {
	sum   int
	terms int
}

type productSumTerm struct {
	m      map[int]map[sumTerm]struct{}
	primes []int
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	psLimit := 2 * limit
	productSumMap := map[int]int{}
	primes := projecteuler.Primes(psLimit, nil)
	pst := productSumTerm{
		m:      map[int]map[sumTerm]struct{}{},
		primes: primes,
	}

	for ps := 4; ps <= psLimit; ps++ {
		var factors map[int]int
		factors, err = projecteuler.Factorize(ps, primes)
		if err != nil {
			return
		}

		var stSet map[sumTerm]struct{}
		stSet, err = pst.factorSumTerm(ps, factors)
		if err != nil {
			return
		}

		for st := range stSet {
			ones := ps - st.sum
			kKey := st.terms + ones
			if kKey > limit || kKey == 1 {
				continue
			}
			minProductSum, present := productSumMap[kKey]
			if !present || ps < minProductSum {
				productSumMap[kKey] = ps
			}
		}
	}

	resultSet := map[int]struct{}{}
	for _, v := range productSumMap {
		resultSet[v] = struct{}{}
	}

	resInt := 0
	for r := range resultSet {
		resInt += r
	}

	result = strconv.Itoa(resInt)
	return
}

func (pst productSumTerm) factorSumTerm(product int, factors map[int]int) (map[sumTerm]struct{}, error) {
	existing, present := pst.m[product]
	if present {
		return existing, nil
	}

	sts := map[sumTerm]struct{}{}
	divisors := projecteuler.FindDivisors(factors)

	if len(divisors) == 2 {
		sts[sumTerm{sum: product, terms: 1}] = struct{}{}
		return sts, nil
	}

	for i := 1; i < len(divisors)-1; i++ { // disregard divisor 1 and product
		rest := product / divisors[i]
		restFactors, err := projecteuler.Factorize(rest, pst.primes)
		if err != nil {
			return sts, err
		}

		restFST, err := pst.factorSumTerm(rest, restFactors)
		if err != nil {
			return sts, err
		}
		restFST[sumTerm{sum: rest, terms: 1}] = struct{}{}

		for k := range restFST {
			newK := sumTerm{
				sum:   k.sum + divisors[i],
				terms: k.terms + 1,
			}
			sts[newK] = struct{}{}
		}
	}

	pst.m[product] = sts
	return sts, nil
}
