package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 16; Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	bi := projecteuler.MakeBigInt("1024")
	bi.PowBigInt(10)
	bi.PowBigInt(10)
	sum := bi.DigitSum()

	result = strconv.Itoa(sum)
	return
}
