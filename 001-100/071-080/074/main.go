package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 74; Digit Factorial Chains
The number 145 is well known for the property that the sum of the factorial of its digits is equal to 145:
1! + 4! + 5! = 1 + 24 + 120 = 145.
Perhaps less well known is 169, in that it produces the longest chain of numbers that link back to 169; it turns out
that there are only three such loops that exist:
169 -> 363601 -> 1454 -> 169
871 -> 45361 -> 871
872 -> 45362 -> 872

It is not difficult to prove that EVERY starting number will eventually get stuck in a loop. For example,
69 -> 363600 -> 1454 -> 169 -> 363601 -> (1454)
78 -> 45360 -> 871 -> 45361 -> (871)
540 -> 145 -> (145)
Starting with 69 produces a chain of five non-repeating terms, but the longest non-repeating chain with a starting number
below one million is sixty terms.
How many chains, with a starting number below one million, contain exactly sixty non-repeating terms?
*/

func main() {
	var limit int

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	digitFactorials := make([]int, 10)
	digitFactorials[0] = 1
	digitFactorials[1] = 1
	for i := 2; i < 10; i++ {
		digitFactorials[i] = digitFactorials[i-1] * i
	}

	maxChainLength := 0
	maxChains := [][]int{}
	for i := 1; i < 1000000; i++ {
		currChain := chain(i, digitFactorials)

		l := len(currChain)
		if l > maxChainLength {
			maxChainLength = l
			maxChains = [][]int{currChain}
		} else if l == maxChainLength {
			maxChains = append(maxChains, currChain)
		}
	}

	// fmt.Printf("Chain length: %d\n", maxChainLength)
	result = strconv.Itoa(len(maxChains))
	return
}

func digitFactorialSum(n int, digitFactorials []int) int {
	sum := 0

	for n > 0 {
		sum += digitFactorials[n%10]
		n /= 10
	}

	return sum
}

func chain(n int, digitFactorials []int) []int {
	chained := []int{n}
	chainedSet := make(map[int]struct{})
	chainedSet[n] = struct{}{}

	for {
		next := digitFactorialSum(n, digitFactorials)
		if _, present := chainedSet[next]; present {
			break
		}

		chained = append(chained, next)
		chainedSet[next] = struct{}{}
		n = next
	}

	return chained
}
