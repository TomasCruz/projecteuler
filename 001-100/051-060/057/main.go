package main

import (
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 57; Square root convergents

It is possible to show that the square root of two can be expressed as an infinite continued fraction.
sqrt(2)=1+1/(2+1/(2+1/(2+...)))

By expanding this for the first four iterations, we get:
						1 + 1/2 = 3/2 = 1.5
				1 + 1/(	2 + 1/2) = 7/5 = 1.4
		1 + 1/(	2 + 1/(	2 + 1/2)) = 17/12 = 1.41666...
1 + 1/(	2 + 1/(	2 + 1/(	2 + 1/2))) = 41/29 = 1.41379...

The next three expansions are 99/70, 239/169, and 577/408, but the eighth expansion, 1393/985,
is the first example where the number of digits in the numerator exceeds the number of digits in the denominator.

In the first one-thousand expansions, how many fractions contain a numerator with more digits than the denominator?
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
		limit = 1000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	prevIter := projecteuler.MakeFraction(big.NewInt(1), big.NewInt(1))
	res := 0
	for i := 1; i < limit; i++ {
		prevIter.AddInt(1)
		prevIter.Invert()
		prevIter.AddInt(1)
		numCount := len(prevIter.Numerator().String())
		denomCount := len(prevIter.Denominator().String())

		if numCount > denomCount {
			res++
		}
	}
	result = strconv.Itoa(res)

	return
}
