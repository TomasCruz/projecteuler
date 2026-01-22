package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 108; Diophantine Reciprocals I
In the following equation x, y, and n are positive integers.

1/x + 1/y = 1/n

For n = 4 there are exactly three distinct solutions:
1/5 + 1/20 = 1/4
1/6 + 1/12 = 1/4
1/8 + 1/8 = 1/4

What is the least value of n for which the number of distinct solutions exceeds one-thousand?

NOTE: This problem is an easier version of Problem 110; it is strongly advised that you solve this one first.
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
		limit = 1000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(1000000, nil)

	n := 4
	for ; ; n++ {
		factors, err := projecteuler.Factorize(n, primes)
		if err != nil {
			panic("can't factorize")
		}

		// square it
		for k, v := range factors {
			factors[k] = 2 * v
		}

		divisors := projecteuler.FindDivisors(factors)

		lesserEqualDivisorsCount := 0
		lesserEqualDivisors := make([]int, 0, len(divisors))
		for _, x := range divisors {
			if x <= n {
				lesserEqualDivisors = append(lesserEqualDivisors, x)
				lesserEqualDivisorsCount++
			}
		}

		if lesserEqualDivisorsCount > limit {
			break
		}
	}

	result = strconv.Itoa(n)
	return
}

/*
	x > n
	for x <= y, 1/x >= 1/y => 1/x + 1/y <= 2/x, 1/n <= 2/x, x <= 2n   !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	x in [n+1,...,2n]
	x-n in [1,...,n]

	1/n = 1/x + 1/y / *nxy
	nx + ny = yx
	nx = y*(x - n)

	y = nx / (x-n) = n*(x-n+n) / (x-n) = n*[1 + n / (x-n)]
	so, divisibility of n^2/1, n^2/2, ... n^2/n

1 2 3 4 5 6 7 8 9 10
 12 14 15 16 18 20 21 24 25 27
 28 30 35 36 40 42 45 48 49 50
 54 56 60 63 70 72 75 80 81 84
 90 98 100 105 108 112 120 126 135 140
 144 147 150 162 168 175 180 189 196 200
 210 216 225 240 245 252 270 280 294 300
 315 324 336 350 360 378 392 400 405 420
 432 441 450 490 504 525 540 560 567 588
 600 630 648 675 700 720 735 756 784 810
 840 882 900 945 980 1008 1050 1080 1134 1176
 1200 1225 1260

 113
*/
