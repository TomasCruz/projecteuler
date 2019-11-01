package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 43; Sub-string divisibility

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order,
but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

    d2d3d4=406 is divisible by 2
    d3d4d5=063 is divisible by 3
    d4d5d6=635 is divisible by 5
    d5d6d7=357 is divisible by 7
    d6d7d8=572 is divisible by 11
    d7d8d9=728 is divisible by 13
    d8d9d10=289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	var res int64
	projecteuler.Permutations(10, sumSubDivisible, &res)
	result = strconv.FormatInt(res, 10)
	return
}

func sumSubDivisible(args ...interface{}) bool {
	resPtr := args[0].(*int64)
	currPerm := args[1].([]byte)

	if currPerm[0] == byte(0) {
		return false
	}

	if int(currPerm[3])%2 != 0 || int(currPerm[5])%5 != 0 {
		return false
	}

	div3 := int(currPerm[2]) + int(currPerm[3]) + int(currPerm[4])
	if div3%3 != 0 {
		return false
	}

	div7 := 100*int(currPerm[4]) + 10*int(currPerm[5]) + int(currPerm[6])
	if div7%7 != 0 {
		return false
	}

	div11 := 100*int(currPerm[5]) + 10*int(currPerm[6]) + int(currPerm[7])
	if div11%11 != 0 {
		return false
	}

	div13 := 100*int(currPerm[6]) + 10*int(currPerm[7]) + int(currPerm[8])
	if div13%13 != 0 {
		return false
	}

	div17 := 100*int(currPerm[7]) + 10*int(currPerm[8]) + int(currPerm[9])
	if div17%17 != 0 {
		return false
	}

	res := 0
	digits := len(currPerm)
	for i := 0; i < digits; i++ {
		res *= 10
		res += int(currPerm[i])
	}

	*resPtr += int64(res)
	return false
}
