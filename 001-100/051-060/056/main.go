package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 56; Powerful digit sum

A googol (10^100) is a massive number: one followed by one-hundred zeros; 100^100 is almost unimaginably large:
one followed by two-hundred zeros. Despite their size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, a^b, where a, b < 100, what is the maximum digital sum?
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
		limit = 100
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	maximumDigitalSum := 0
	for i := 1; i < limit; i++ {
		for j := 1; j < limit; j++ {
			num := projecteuler.MakeBigIntFromInt(i)
			num.PowBigInt(j)
			ds := num.DigitSum()
			if ds > maximumDigitalSum {
				maximumDigitalSum = ds
			}
		}
	}
	result = strconv.Itoa(maximumDigitalSum)

	return
}
