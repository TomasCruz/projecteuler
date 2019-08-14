package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 33; Digit cancelling fractions

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may
incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value,
and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	prodNumer, prodDenom := 1, 1

	for denom := 11; denom < 100; denom++ {
		if denom%10 == 0 {
			continue
		}

		for numer := 10; numer < denom; numer++ {
			if numer%10 == 0 {
				continue
			}

			n := projecteuler.NewDigitalNumber(numer)
			d := projecteuler.NewDigitalNumber(denom)

			nDigits := n.Digits()
			dDigits := d.Digits()
			if nDigits[0] == 0 || nDigits[1] == 0 || dDigits[0] == 0 || dDigits[1] == 0 {
				continue
			}

			if (nDigits[0] == dDigits[1] && int(nDigits[1])*denom == int(dDigits[0])*numer) ||
				(nDigits[1] == dDigits[0] && int(nDigits[0])*denom == int(dDigits[1])*numer) {

				prodNumer *= numer
				prodDenom *= denom
			}
		}
	}

	primes := projecteuler.Primes(100, nil) // two digit numbers can't have prime factors larger than 100
	fNum, _ := projecteuler.Factorize(prodNumer, primes)
	fDen, _ := projecteuler.Factorize(prodDenom, primes)

	fResDen := make(map[int]int)
	for f, dv := range fDen {
		nv := fNum[f]

		min := nv
		if dv < nv {
			min = dv
		}

		val := dv - min
		if val != 0 {
			fResDen[f] = val
		}
	}

	denomValue := projecteuler.MultiplyFactors(fResDen)
	result = strconv.FormatInt(denomValue, 10)

	return
}
