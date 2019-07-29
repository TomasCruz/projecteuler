package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 14; Longest Collatz sequence

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
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

	projecteuler.TimedStr(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	resultInt := 0
	resultLength := 0
	for i := 2; i < limit; i++ {
		curr := collatzLength(i)
		if curr > resultLength {
			resultInt = i
			resultLength = curr
		}
	}

	result = strconv.Itoa(resultInt)
	return
}

func collatzLength(num int) (length int) {
	length = 1

	for {
		if num%2 == 0 {
			num /= 2
		} else {
			num = 3*num + 1
		}

		length++
		if num == 1 {
			break
		}
	}

	return
}
