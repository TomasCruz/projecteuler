package main

import (
	"fmt"

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

func calc(args ...interface{}) (err error) {
	var bi projecteuler.BigInt
	if bi, err = projecteuler.MakeBigInt("1024"); err != nil {
		return
	}

	biPtr := &bi
	biPtr.PowBigInt(10)
	biPtr.PowBigInt(10)
	sum := bi.DigitSum()
	fmt.Println(sum)

	return
}
