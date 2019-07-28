package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 15; Lattice paths

Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down,
there are exactly 6 routes to the bottom right corner.

rrdd, rdrd, rddr, drrd, drdr, ddrr
How many such routes are there through a 20×20 grid?
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
		limit = 20
	}

	projecteuler.TimedStr(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(2*limit+1, nil)

	var binomial map[int]int
	if binomial, err = projecteuler.Binomial(2*limit, limit, primes); err != nil {
		return
	}

	resultInt64 := projecteuler.MultiplyFactors(binomial)

	result = strconv.FormatInt(resultInt64, 10)
	return
}
