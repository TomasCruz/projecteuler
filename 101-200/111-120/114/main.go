package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 114; Counting Block Combinations I
A row measuring seven units in length has red blocks with a minimum length of three units placed on it,
such that any two red blocks (which are allowed to be different lengths) are separated by at least one grey square.
There are exactly seventeen ways of doing this.

How many ways can a row measuring fifty units in length be filled?

NOTE: Although the example above does not lend itself to the possibility, in general it is permitted to mix block sizes.
For example, on a row measuring eight units in length you could use red (3), grey (1), and red (4).
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

	m := [5][]uint64{}
	for i := 0; i < 5; i++ {
		m[i] = make([]uint64, limit+1)
	}

	m[0][3] = 1
	m[0][4] = 2
	m[1][1] = 1
	m[1][4] = 1
	m[2][2] = 1
	m[3][3] = 1
	m[4][4] = 1

	for i := 5; i <= limit; i++ {
		m[0][i] = m[0][i-1] + m[3][i-1] + m[4][i-1]
		m[1][i] = m[0][i-1]
		m[2][i] = m[1][i-1]
		m[3][i] = m[2][i-1]
		m[4][i] = m[3][i-1] + m[4][i-1]
	}

	sum := m[0][limit] + m[1][limit] + m[2][limit] + m[3][limit] + m[4][limit]
	result = strconv.FormatUint(sum, 10)
	return
}
