package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 34; Digit factorials

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	digitFact := make([]int, 10)
	digitFact[0] = 1
	for i := 1; i < 10; i++ {
		digitFact[i] = i * digitFact[i-1]
	}

	resultSum := 0
	limit := 7 * digitFact[9]
	for x := 10; x <= limit; x++ {
		dig := projecteuler.NewDigitalNumber(x)
		digits := dig.Digits()
		digSum, digCount := 0, len(digits)

		for i := 0; i < digCount; i++ {
			digSum += digitFact[int(digits[i])]
		}

		if x == digSum {
			resultSum += x
		}
	}

	result = strconv.Itoa(resultSum)
	return
}
