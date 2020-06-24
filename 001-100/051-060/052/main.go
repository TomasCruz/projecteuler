package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 52; Permuted multiples

It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits,
but in a different order.
Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
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
		limit = 1000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	for i := 1; i < limit; i++ {
		if checkNumber(i) {
			result = strconv.FormatInt(int64(i), 10)
			break
		}
	}

	return
}

func checkNumber(x int) bool {
	dn := projecteuler.NewDigitalNumber(x)

	d2 := projecteuler.NewDigitalNumber(2 * x)
	if !dn.SameDigitSet(d2) {
		return false
	}

	d3 := projecteuler.NewDigitalNumber(3 * x)
	if !dn.SameDigitSet(d3) {
		return false
	}

	d4 := projecteuler.NewDigitalNumber(4 * x)
	if !dn.SameDigitSet(d4) {
		return false
	}

	d5 := projecteuler.NewDigitalNumber(5 * x)
	if !dn.SameDigitSet(d5) {
		return false
	}

	d6 := projecteuler.NewDigitalNumber(6 * x)
	if !dn.SameDigitSet(d6) {
		return false
	}

	return true
}
