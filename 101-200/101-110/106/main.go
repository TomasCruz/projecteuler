package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 106; Special Subset Sums: Meta-testing
Let S(A) represent the sum of elements in set A of size n. We shall call it a special sum set if for any two non-empty disjoint subsets,
B and C, the following properties are true:
	S(B) != S(C); that is, sums of subsets cannot be equal.
	If B contains more elements than C then S(B) > S(C).

For this problem we shall assume that a given set contains n strictly increasing elements and it already satisfies the second rule.

Surprisingly, out of the 25 possible subset pairs that can be obtained from a set for which n = 4, only 1 of these pairs need to be tested for equality (first rule).
Similarly, when n = 7, only 70 out of the 966 subset pairs need to be tested.

For n = 12, how many of the 261625 subset pairs that can be obtained need to be tested for equality?

NOTE: This problem is related to Problem 103 and Problem 106.
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
		limit = 12
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	half := limit / 2
	sum := 0
	for i := 2; i <= half; i++ {
		currX := 0
		projecteuler.Combinations(byte(2*i), byte(i), x, &currX)
		sum += numCombos(limit, 2*i) * currX
	}

	result = strconv.Itoa(sum)
	return
}

func x(args ...interface{}) bool {
	resPtr := args[0].(*int)
	combo := args[1].([]byte)

	if combo[0] == 1 {
		return false
	}

	l := len(combo) / 2
	zeroes := make([]byte, 0, l)
	ones := make([]byte, 0, l)

	for i := range combo {
		if combo[i] == 0 {
			zeroes = append(zeroes, byte(i))
		} else {
			ones = append(ones, byte(i))
		}
	}

	for i := 0; i < l; i++ {
		if zeroes[i] > ones[i] {
			*resPtr++
			break
		}
	}

	return false
}

func numCombos(n, k int) int {
	up := 1
	down := 1

	for i := 0; i < k; i++ {
		up *= n - i
		down *= k - i
	}

	return up / down
}
