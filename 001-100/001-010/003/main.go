package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Largest prime factor
Problem 3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

func main() {
	var limit int64

	if len(os.Args) > 1 {
		var err error
		limit, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}
	} else {
		limit = 600851475143
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int64)

	var largest int64
	inspected := int64(limit)
	projecteuler.PrimesDivisibility(inspected/2+1, isDivisibleToOne, &inspected, &largest)

	if largest == 0 {
		largest = int64(limit)
	}

	result = strconv.FormatInt(largest, 10)
	return
}

func isDivisibleToOne(args ...interface{}) bool {
	inspected := args[0].(*int64)
	largest := args[1].(*int64)
	divider := args[2].(int64)

	for *inspected%divider == 0 {
		*largest = divider
		*inspected /= divider
	}

	return *inspected == 1
}
