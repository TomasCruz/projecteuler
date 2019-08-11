package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 32; Pandigital products

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once;
for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier,
and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	/*
		For x with a digits, 10^(a-1) <= x < 10^a
		For y with b digits, 10^(b-1) <= y < 10^b
		10^(a+b-2) <= x*y < 10^(a+b) => x*y has a+b-1 or a+b digits
		So, x,y,xy have 2*(a+b) or 2*(a+b)-1 digits. As this number should be 9, it must be 2*(a+b)-1
		as the other possibility is even number. So x and y together have 5 digits, and xy has 4 digits.
		Smaller of these two factors has either 1 digit and greater 4, or smaller has 2 digits and greater 3.
	*/

	primes := projecteuler.Primes(10000, nil)
	pandigitalProducts := make(map[int]struct{})

	for product := 1234; product <= 9876; product++ {
		dn := projecteuler.NewDigitalNumber(product)
		if differentNonZeroDigits(dn) && checkFactors(dn, primes) {
			pandigitalProducts[product] = struct{}{}
		}
	}

	sum := 0
	for k := range pandigitalProducts {
		sum += k
	}

	result = strconv.Itoa(sum)
	return
}

func differentNonZeroDigits(dn projecteuler.DigitalNumber) bool {
	if _, ok := dn.DigitMap()[0]; ok {
		return false
	}

	return dn.DigitCount() == len(dn.DigitMap())
}

func checkFactors(dn projecteuler.DigitalNumber, primes []int) bool {
	factors, _ := projecteuler.Factorize(dn.X(), primes)
	divisors := projecteuler.FindDivisors(factors)

	for _, d := range divisors {
		if d == 1 || d > 99 {
			continue
		}

		smallerNd := projecteuler.NewDigitalNumber(d)
		if !differentNonZeroDigits(smallerNd) {
			continue
		}

		if differentDigits, usedDigits := differentDigitCompositions(dn, smallerNd); differentDigits {
			noPandigital := false
			rest := dn.X() / d

			for rest > 0 {
				if _, ok := usedDigits[byte(rest%10)]; ok {
					noPandigital = true // no pandigital property
					break
				}

				rest /= 10
			}

			if !noPandigital {
				greaterNd := projecteuler.NewDigitalNumber(dn.X() / d)
				if !differentNonZeroDigits(greaterNd) {
					continue
				}

				if greaterNd.DigitCount()+len(usedDigits) == 9 {
					return true
				}
			}
		}
	}

	return false
}

func differentDigitCompositions(a, b projecteuler.DigitalNumber) (differentDigits bool, usedDigits map[byte]struct{}) {
	usedDigits = make(map[byte]struct{})

	for _, k := range a.Digits() {
		usedDigits[k] = struct{}{}
	}

	for _, k := range b.Digits() {
		if _, ok := usedDigits[k]; ok {
			return
		}

		usedDigits[k] = struct{}{}
	}

	differentDigits = true
	return
}
