package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 26; Reciprocal cycles

A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions
with denominators 2 to 10 are given:

    1/2	= 	0.5
    1/3	= 	0.(3)
    1/4	= 	0.25
    1/5	= 	0.2
    1/6	= 	0.1(6)
    1/7	= 	0.(142857)
    1/8	= 	0.125
    1/9	= 	0.(1)
    1/10	= 	0.1

Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
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

	maxRecurringLength := 0
	maxDenom := 0
	for i := 2; i < limit; i++ {
		currRecurring := findRecurringCycleLength(i)
		if currRecurring > maxRecurringLength {
			maxRecurringLength = currRecurring
			maxDenom = i
		}
	}

	result = strconv.Itoa(maxDenom)
	return
}

func findRecurringCycleLength(denom int) int {
	nominMap := make(map[int]int)
	return rec(1, denom, 0, nominMap)
}

func rec(nomin, denom, index int, nominMap map[int]int) (length int) {
	nomin *= 10
	if nomin%denom == 0 {
		return 0
	}

	if prev, ok := nominMap[nomin]; ok {
		return index - prev
	}

	nominMap[nomin] = index
	return rec(nomin%denom, denom, index+1, nominMap)
}
