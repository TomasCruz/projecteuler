package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 159; Digital Root Sums of Factorisations
A composite number can be factored many different ways.
For instance, not including multiplication by one, 24 can be factored in 7 distinct ways:
24 = 2 * 2 * 2 * 3
24 = 2 * 3 * 4
24 = 2 * 2 * 6
24 = 4 * 6
24 = 3 * 8
24 = 2 * 12
24 = 24

Recall that the digital root of a number, in base 10, is found by adding together the digits of that number,
and repeating that process until a number is arrived at that is less than 10.
Thus the digital root of 467 is 8.

We shall call a Digital Root Sum (DRS) the sum of the digital roots of the individual factors of our number.
The chart below demonstrates all of the DRS values for 24.
Factorisation		Digital Root Sum
2 * 2 * 2 * 3		9
2 * 3 * 4			9
2 * 2 * 6			10
4 * 6				10
3 * 8				11
2 * 12				5
24					6

The maximum Digital Root Sum of 24 is 11.
The function mdrs(n) gives the maximum Digital Root Sum of n. So mdrs(24)=11.

Find sum mdrs(n) for 1 < n < 1,000,000.
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

	mdrs := make([]int, limit)

	rt := int(math.Sqrt(float64(limit)))
	for i := 2; i < limit; i++ {
		mdrs[i] = dr(i)
	}

	for i := 2; i <= rt; i++ {
		for j := 1; i*j < limit; j++ {
			mdrs[i*j] = max(mdrs[i*j], mdrs[i]+mdrs[j])
		}
	}

	sum := 0
	for i := 2; i < limit; i++ {
		sum += mdrs[i]
	}

	result = strconv.Itoa(sum)
	return
}

func drs(factorization []int) int {
	sum := 0
	for _, r := range factorization {
		sum += dr(r)
	}

	return sum
}

func dr(x int) int {
	ret := x % 9
	if ret == 0 {
		ret = 9
	}

	return ret
}

/*
	drs is by definition multiplicative for a specific factorization.
	For each particular factorization n=a*b, mdrs(n) = max(dr(n), mdrs(a)+mdrs(b))
	That implies that for each potential factor, it is needed to set current value of mdrs
	to maximum of it's previous iteration and sum of it's factors' previous iteration of mdrs.
	Solution is a modified sieve of Erathostenes algorithm that will update array as neccessary.
*/
