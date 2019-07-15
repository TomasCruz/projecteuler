package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 10; Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
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
		limit = 2000000
	}

	args = append(args, limit)
	projecteuler.Timed(calc, args...)
}

func calc(args ...interface{}) (err error) {
	limit := args[0].(int)
	primes := projecteuler.Primes(limit, nil)

	var result int64
	for _, x := range primes {
		result += int64(x)
	}

	fmt.Println(result)
	return
}
