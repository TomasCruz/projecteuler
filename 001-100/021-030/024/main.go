package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 24; Lexicographic permutations

A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation
of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it
lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

func main() {
	var limit byte
	var ordinal int

	if len(os.Args) == 3 {
		limit64, err := strconv.ParseInt(os.Args[1], 10, 8)
		if err != nil {
			log.Fatal("bad argument")
		}
		limit = byte(limit64)

		ordinal64, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}
		ordinal = int(ordinal64)
	} else {
		limit = 10
		ordinal = 1000000
	}

	projecteuler.Timed(calc, limit, ordinal)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(byte)
	ordinal := args[1].(int)
	resultPerm := make([]byte, limit)

	projecteuler.Permutations(limit, func(args ...interface{}) bool {
		ordinal := args[0].(*int)
		resultPerm := args[1].([]byte)
		currPerm := args[2].([]byte)

		*ordinal--
		if *ordinal == 0 {
			copy(resultPerm, currPerm)
			return true
		}

		return false
	}, &ordinal, resultPerm)

	for i := byte(0); i < limit; i++ {
		resultPerm[i] += '0'
	}

	result = string(resultPerm)
	return
}
