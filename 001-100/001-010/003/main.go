package main

import (
	"fmt"
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
	var args []interface{}

	if len(os.Args) > 1 {
		limit, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		args = append(args, limit)
	}

	projecteuler.Timed(calc, args...)
}

func calc(args ...interface{}) {
	var limit int64

	if len(args) == 0 {
		limit = 600851475143
	} else {
		limit = args[0].(int64)
	}

	var largest int
	inspected := int(limit)
	projecteuler.Primes(inspected/2+1, isDivisibleToOne, &inspected, &largest)

	if largest == 0 {
		largest = int(limit)
	}

	fmt.Println(largest)
}

func isDivisibleToOne(args ...interface{}) bool {
	inspected := args[0].(*int)
	largest := args[1].(*int)
	divider := args[2].(int)

	for *inspected%divider == 0 {
		*largest = divider
		*inspected /= divider
	}

	return *inspected == 1
}
