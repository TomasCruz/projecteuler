package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 135; Same Differences
Given the positive integers, x, y, and z, are consecutive terms of an arithmetic progression,
the least value of the positive integer, n, for which the equation, x^2 - y^2 - z^2 = n, has exactly two solutions is n = 27:
34^2 - 27^2 - 20^2 = 12^2 - 9^2 - 6^2 = 27.

It turns out that n = 1155 is the least value which has exactly ten solutions.
How many values of n less than one million have exactly ten distinct solutions?
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
		if ns == 10 {
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
