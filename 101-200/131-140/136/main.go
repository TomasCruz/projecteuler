package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 136; Singleton Difference
The positive integers, x, y, and z, are consecutive terms of an arithmetic progression. Given that n is a positive integer,
the equation, x^2 - y^2 - z^2 = n, has exactly one solution when n = 20:
13^2 - 10^2 - 7^2 = 20.

In fact there are twenty-five values of n below one hundred for which the equation has a unique solution.
How many values of n less than fifty million have exactly one solution?
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
		limit = 50000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	numSolutions := make([]int, limit)
	for y := 1; y < limit; y++ {
		for d := y/4 + 1; d < y; d++ {
			n := y * (4*d - y)
			if n >= limit {
				break
			}

			numSolutions[n]++
		}
	}

	count := 0
	for _, ns := range numSolutions {
		if ns == 1 {
			count++
		}
	}

	result = strconv.Itoa(count)
	return
}

/*
	(y + d)^2 - y^2 - (y - d)^2 = n, y > d
	4*y*d - y^2 = n
	y*(4*d - y) = n
	4*d - y > 0
	d > y/4
	y/4 < d < y
*/
