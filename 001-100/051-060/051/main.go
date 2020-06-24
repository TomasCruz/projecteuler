package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 51; Prime digit replacements

By replacing the 1st digit of the 2-digit number *3, it turns out that six of the nine possible values:
13, 23, 43, 53, 73, and 83, are all prime.
By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit number is the first example
having seven primes among the ten generated numbers, yielding the family:
56003, 56113, 56333, 56443, 56663, 56773, and 56993. Consequently 56003, being the first member of this family,
is the smallest prime with this property.

Find the smallest prime which, by replacing part of the number (not necessarily adjacent digits)
with the same digit, is part of an eight prime value family.
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

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes, primeSet := projecteuler.PrimeSet(limit)

	for _, p := range primes {
		dn := projecteuler.NewDigitalNumber(p)
		digitCount := dn.DigitCount()

		for digitsToReplace := 1; digitsToReplace < digitCount; digitsToReplace++ {
			combinations, _ := projecteuler.Combinations(byte(digitCount), byte(digitsToReplace), nil)
			for _, c := range combinations {
				numPrimes := 10
				replacementDigitStart := 0
				if c[0] == 1 {
					replacementDigitStart++
					numPrimes--
				}

				smallest := 0
				for replacementDigit := replacementDigitStart; replacementDigit < 10; replacementDigit++ {
					replacements := make(map[byte]byte)
					for d := 0; d < digitCount; d++ {
						if c[d] == 1 {
							replacements[byte(d)] = byte(replacementDigit)
						}
					}

					trueComb := dn.ReplaceDigits(replacements)
					if _, ok := primeSet[trueComb]; ok {
						if replacementDigit == replacementDigitStart {
							smallest = trueComb
						}
					} else {
						numPrimes--
						if numPrimes < 8 {
							break
						}
					}
				}

				if numPrimes == 8 {
					result = strconv.FormatInt(int64(smallest), 10)
					return
				}
			}
		}
	}

	return
}
