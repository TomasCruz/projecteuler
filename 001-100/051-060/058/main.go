package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 58; Spiral primes

Starting with 1 and spiralling anticlockwise in the following way, a square spiral with side length 7 is formed.

37 36 35 34 33 32 31
38 17 16 15 14 13 30
39 18  5  4  3 12 29
40 19  6  1  2 11 28
41 20  7  8  9 10 27
42 21 22 23 24 25 26
43 44 45 46 47 48 49

It is interesting to note that the odd squares lie along the bottom right diagonal, but what is more interesting is
that 8 out of the 13 numbers lying along both diagonals are prime; that is, a ratio of 8/13 ≈ 62%.

If one complete new layer is wrapped around the spiral above, a square spiral with side length 9 will be formed.
If this process is continued, what is the side length of the square spiral for which the ratio of primes along both
diagonals first falls below 10%?
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
		limit = 15000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	diagonalPrimeCount := 0
	diagonalCount := 1

	corners := make([]int, 4)
	for j := 0; j < 4; j++ {
		corners[j] = 1
	}

	for i := 0; i < limit; i++ {
		corners[0] += 8*i + 2 // ur
		corners[1] += 8*i + 4 // ul
		corners[2] += 8*i + 6 // dl
		corners[3] += 8*i + 8 // dr

		for j := 0; j < 4; j++ {
			if projecteuler.IsPrime(int64(corners[j])) {
				diagonalPrimeCount++
			}
		}

		diagonalCount += 4
		ratio := float64(diagonalPrimeCount) / float64(diagonalCount)
		if ratio < 0.1 {
			result = strconv.Itoa(2*i + 3)
			return
		}
	}

	result = "-1"
	return
}
