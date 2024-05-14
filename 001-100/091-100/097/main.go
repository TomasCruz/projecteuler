package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 97; Large Non-Mersenne Prime
The first known prime found to exceed one million digits was discovered in 1999, and is a Mersenne prime of the form 2^6972593 - 1;
it contains exactly 2,098,960 digits. Subsequently other Mersenne primes, of the form 2^p - 1, have been found which contain more digits.

However, in 2004 there was found a massive non-Mersenne prime which contains 2,357,207 digits: 28433 * 2^7830457 + 1.

Find the last ten digits of this prime number.
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	// the huge number's last 10 digits, obtained through multiplication and addition,
	// depends only of the last 10 digits of each step of the calculation.

	degree := projecteuler.MakeBigIntFromInt(2)
	for i := 1; i < 7830457; i++ {
		degree.MultiplyByDigit(2)
		if degree.DigitCount() > 10 {
			s := degree.String()
			s = s[len(s)-10:]
			degree = projecteuler.MakeBigInt(s)
		}
	}
	res := degree.Int()
	res *= 28433
	res += 1

	result = strconv.FormatInt(res, 10)
	result = result[len(result)-10:]

	return
}
