package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 95; Amicable Chains
The proper divisors of a number are all the divisors excluding the number itself. For example, the proper divisors of 28 are 1, 2, 4, 7, and 14.
As the sum of these divisors is equal to 28, we call it a perfect number.

Interestingly the sum of the proper divisors of 220 is 284 and the sum of the proper divisors of 284 is 220, forming a chain of two numbers.
For this reason, 220 and 284 are called an amicable pair.

Perhaps less well known are longer chains. For example, starting with 12496, we form a chain of five numbers:
12496 -> 14288 -> 15472 -> 14536 -> 14264 ( -> 12496 -> ...)

Since this chain returns to its starting point, it is called an amicable chain.
Find the smallest member of the longest amicable chain with no element exceeding one million.
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
		limit = 1000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	primes := projecteuler.Primes(limit, nil)
	pd := properDivisors{
		limit:  limit,
		primes: primes,
	}

	pd.buildPrimePowers()
	pd.buildPrimeDivisorSum()

	properDivisorsSum := make([]int, limit+1)
	for i := 2; i <= limit; i++ {
		p := divisorSum(i, primes, pd.primeDivisorSum) - i

		if p > limit {
			properDivisorsSum[i] = 0
		} else {
			properDivisorsSum[i] = p
		}
	}

	smallestChainStart := 0
	maxChainLength := 0
	for i := 2; i <= limit; i++ {
		currChainLength := chainLength(i, limit, properDivisorsSum)
		if currChainLength > maxChainLength {
			smallestChainStart = i
			maxChainLength = currChainLength
		}
	}

	result = strconv.Itoa(smallestChainStart)
	return
}

func divisorSum(x int, primes []int, primeDivisorSum map[int][]int64) int {
	factors, err := projecteuler.Factorize(x, primes)
	if err != nil {
		return 0
	}

	p := int64(1)
	for k, v := range factors {
		p *= primeDivisorSum[k][v]
	}

	return int(p)
}

func chainLength(i, limit int, properDivisorsSum []int) int {
	chain := map[int]struct{}{}
	prev := i
	next := 0

	for {
		next = properDivisorsSum[prev]

		if next == 0 || next > limit {
			return 0
		}

		if _, exists := chain[prev]; exists {
			break
		}

		chain[prev] = struct{}{}
		if next == i {
			return len(chain)
		}

		prev = next
	}

	return 0
}

type properDivisors struct {
	limit           int
	primes          []int
	primePowers     [][]int64
	primeDivisorSum map[int][]int64
}

func (pd *properDivisors) buildPrimePowers() {
	l := len(pd.primes)
	pd.primePowers = make([][]int64, l)

	for i := 0; i < l; i++ {
		pd.primePowers[i] = []int64{int64(1), int64(pd.primes[i])}
		for j := 2; ; j++ {
			pd.primePowers[i] = append(pd.primePowers[i], pd.primePowers[i][j-1]*int64(pd.primes[i]))
			if pd.primePowers[i][j] > int64(pd.limit) {
				break
			}
		}
	}
}

// https://mathworld.wolfram.com/DivisorFunction.html
func (pd *properDivisors) buildPrimeDivisorSum() {
	l := len(pd.primes)
	pd.primeDivisorSum = make(map[int][]int64, l)

	for i := 0; i < l; i++ {
		pd.primeDivisorSum[pd.primes[i]] = []int64{int64(0)}
		for j := 1; j < len(pd.primePowers[i])-1; j++ {
			divSum := (pd.primePowers[i][j+1] - int64(1)) / (pd.primePowers[i][1] - int64(1))
			pd.primeDivisorSum[pd.primes[i]] = append(pd.primeDivisorSum[pd.primes[i]], divSum)
		}
	}
}
