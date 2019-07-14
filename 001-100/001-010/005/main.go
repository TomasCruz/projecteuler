package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 5; Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
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
		limit = 20
	}

	args = append(args, limit)
	projecteuler.Timed(calc, args...)
}

func calc(args ...interface{}) (err error) {
	limit := args[0].(int)

	primeFactors := make(map[int]*projecteuler.Powered)
	primes := projecteuler.Primes(limit+1, nil)

	for i := 2; i <= limit; i++ {
		var factors []projecteuler.Powered
		if factors, err = projecteuler.Factorize(i, primes); err != nil {
			fmt.Println(err)
			return
		}

		for _, currFactor := range factors {
			if powered, ok := primeFactors[currFactor.Base]; ok {
				if currFactor.Exp > powered.Exp {
					powered.Exp = currFactor.Exp
				}
			} else {
				primeFactors[currFactor.Base] = &currFactor
			}
		}
	}

	result := multiplyPrimeFactors(primeFactors)
	fmt.Println(result)

	return
}

func multiplyPrimeFactors(primeFactors map[int]*projecteuler.Powered) int64 {
	result := int64(1)

	for _, v := range primeFactors {
		for j := 0; j < v.Exp; j++ {
			result *= int64(v.Base)
		}
	}

	return result
}
