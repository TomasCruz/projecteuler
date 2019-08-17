package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 38; Pandigital multiples

Take the number 192 and multiply it by each of 1, 2, and 3:

    192 × 1 = 192
    192 × 2 = 384
    192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576
the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645,
which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product
of an integer with (1,2, ... , n) where n > 1?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	// smallest number to form pandigital product is 1 with (1,2,3,4,5,6,7,8,9)
	// biggest is when n==2, i.e. a number which multiplied by one produces 4 digit and by two 5 digits,
	// so that number has 4 digits max

	largestProduct := 0
	for i := 1; i < 10000; i++ {
		if pandigital, product := checkNumber(i); pandigital && product > largestProduct {
			largestProduct = product
		}
	}

	result = strconv.Itoa(largestProduct)
	return
}

func checkNumber(x int) (pandigital bool, product int) {
	product = 0

	usedDigits := make(map[byte]struct{})
	for i := 1; i < 10; i++ {
		if len(usedDigits) == 9 {
			pandigital = true
			return
		}

		nextInt := x * i
		nextProduct := projecteuler.NewDigitalNumber(nextInt)
		if !nextProduct.NonZeroDigits() || !nextProduct.DifferentDigits() {
			return
		}

		if differentDigits := nextProduct.DifferentDigitCompositions(usedDigits); !differentDigits {
			return
		}

		product = concatNumbers(product, nextInt)
	}

	return
}

func concatNumbers(lhs, rhs int) int {
	r := projecteuler.NewDigitalNumber(rhs)
	rDigites := r.Digits()

	for i := 0; i < len(rDigites); i++ {
		lhs *= 10
	}

	return lhs + rhs
}
