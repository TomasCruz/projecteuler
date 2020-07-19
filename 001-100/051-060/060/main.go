package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 60; Prime pair sets

The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and concatenating them
in any order the result will always be prime. For example, taking 7 and 109, both 7109 and 1097 are prime.
The sum of these four primes, 792, represents the lowest sum for a set of four primes with this property.

Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.
*/

func main() {
	var limit, setSize int

	if len(os.Args) > 1 {
		setSize64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		setSize = int(setSize64)

		limit64, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limit = int(limit64)
	} else {
		setSize = 5
		limit = 1000000
	}

	projecteuler.Timed(calc, setSize, limit)
}

func calc(args ...interface{}) (result string, err error) {
	setSize := args[0].(int)
	limit := args[1].(int)

	kbs := newKidBrothers(setSize, limit)

	kbs.findLowestSetSum()
	if kbs.winnerSum == math.MaxInt32 {
		result = "-1"
	} else {
		result = strconv.Itoa(kbs.winnerSum)
	}

	return
}
