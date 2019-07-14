package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 7; 10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

func main() {
	var args []interface{}
	var limit int

	if len(os.Args) > 1 {
		limit64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limit = int(limit64)
	} else {
		limit = 10001
	}

	args = append(args, limit)
	projecteuler.Timed(calc, args...)
}

func calc(args ...interface{}) (err error) {
	limit := args[0].(int)

	projecteuler.Primes(math.MaxInt64, isAskedIndex, limit)
	fmt.Println(x)

	return
}

var x int
var ordinal = 2

func isAskedIndex(args ...interface{}) bool {
	limit := args[0].(int)
	currPrime := args[1].(int)

	ordinal++
	if ordinal == limit {
		x = currPrime
		return true
	}

	return false
}
