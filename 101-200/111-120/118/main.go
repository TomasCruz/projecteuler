package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 118; Pandigital Prime Sets
Using all of the digits 1 through 9 and concatenating them freely to form decimal integers, different sets can be formed.
Interestingly with the set {2,5,47,89,631}, all of the elements belonging to it are prime.

How many distinct sets containing each of the digits one through nine exactly once contain only prime elements?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	primes, primeSet := projecteuler.PrimeSet(1000000)
	sets := map[string]struct{}{}

	projecteuler.Permutations(9, func(args ...interface{}) bool {
		primes := args[0].([]int)
		primeSet := args[1].(map[int]struct{})
		sets := args[2].(map[string]struct{})
		currPerm := args[3].([]byte)

		cp := make([]int, 9)
		for i := range currPerm {
			cp[i] = int(currPerm[i] + 1)
		}

		if cp[8]%2 == 1 && cp[8] != 5 {
			findSets(cp, primes, primeSet, sets)
		}

		return false
	}, primes, primeSet, sets)

	result = strconv.Itoa(len(sets))
	return
}

func findSets(currPerm, primes []int, primeSet map[int]struct{}, sets map[string]struct{}) {
	findSetsRec(0, []int{}, currPerm, primes, primeSet, sets)
}

func findSetsRec(start int, primesSoFar, currPerm []int, primes []int, primeSet map[int]struct{}, sets map[string]struct{}) {
	if start == 9 {
		sets[intSliceToString(primesSoFar)] = struct{}{}
		return
	}

	for end := start + 1; end < 10; end++ {
		num := digitSliceToInt(currPerm[start:end])

		if end-start <= 6 {
			if _, present := primeSet[num]; !present {
				continue
			}
		} else {
			if !isPrime(num, primes) {
				continue
			}
		}

		newPrimesSoFar := make([]int, len(primesSoFar)+1)
		copy(newPrimesSoFar, primesSoFar)
		newPrimesSoFar[len(primesSoFar)] = num
		findSetsRec(end, newPrimesSoFar, currPerm, primes, primeSet, sets)
	}
}

func digitSliceToInt(digs []int) int {
	num := 0
	for i := 0; i < len(digs); i++ {
		num *= 10
		num += digs[i]
	}
	return num
}

func intSliceToString(primesSoFar []int) string {
	sort.Ints(primesSoFar)

	s := fmt.Sprintf("%d", primesSoFar[0])
	for i := 1; i < len(primesSoFar); i++ {
		s = fmt.Sprintf("%s,%d", s, primesSoFar[i])
	}

	return s
}

func isPrime(x int, primes []int) bool {
	for _, p := range primes {
		if p*p > x {
			break
		}

		if x%p == 0 {
			return false
		}
	}

	return true
}
