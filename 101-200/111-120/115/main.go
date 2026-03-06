package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 115; Counting Block Combinations II
NOTE: This is a more difficult version of Problem 114.

A row measuring n units in length has red blocks with a minimum length of m units placed on it,
such that any two red blocks (which are allowed to be different lengths) are separated by at least one black square.

Let the fill-count function, F(m, n), represent the number of ways that a row can be filled.

For example, F(3, 29) = 673135 and F(3, 30) = 1089155.
That is, for m = 3, it can be seen that n = 30 is the smallest value for which the fill-count function first exceeds one million.

In the same way, for m = 10, it can be verified that F(10, 56) = 880711 and F(10, 57) = 1148904,
so n = 57 is the least value for which the fill-count function first exceeds one million.

For m = 50, find the least value of n for which the fill-count function first exceeds one million.
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
		limit = 50
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	m := make([][]uint64, limit+3)
	for i := 0; i < limit+3; i++ {
		m[i] = make([]uint64, 1000)
	}

	for i := 1; i <= limit+1; i++ {
		m[i][i] = 1
	}
	m[0][limit] = 1
	m[0][limit+1] = 2
	m[1][limit+1] = 1

	i := limit + 2
	for ; ; i++ {
		m[0][i] = m[0][i-1]
		for j := limit; j < limit+2; j++ {
			m[0][i] += m[j][i-1]
		}

		for j := 1; j <= limit; j++ {
			m[j][i] = m[j-1][i-1]
		}

		m[limit+1][i] = m[limit][i-1] + m[limit+1][i-1]

		m[limit+2][i] = m[0][i]
		for j := 1; j < limit+2; j++ {
			m[limit+2][i] += m[j][i]
		}

		if m[limit+2][i] > 1000000 {
			break
		}
	}

	result = strconv.Itoa(i)
	return
}
