package main

import (
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 124; Ordered Radicals
The radical of n, rad(n), is the product of the distinct prime factors of n. For example, 504 = 2^3 * 3^2 * 7, so rad(504) = 2 * 3 * 7 = 42.
If we calculate rad(n) for 1 <= n <= 10, then sort them on rad(n), and sorting on n if the radical values are equal, we get:

Unsorted		Sorted
n	rad(n)		n	rad(n)	k
1	1			1		1		1
2	2			2		2		2
3	3			4		2		3
4	2			8		2		4
5	5			3		3		5
6	6			9		3		6
7	7			5		5		7
8	2			6		6		8
9	3			7		7		9
10	10			10		10		10

Let E(k) be the k-th element in the sorted n column; for example, E(4) = 8 and E(6) = 9.
If rad(n) is sorted for 1 <= n <= 100000, find E(10000).
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
		limit = 100000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(limit, nil)

	rds := make([]rd, limit+1)
	for i := 1; i <= limit; i++ {
		factors, _ := projecteuler.Factorize(i, primes)
		currRad := rad(factors)
		rds[i] = rd{n: i, rad: currRad}
	}

	slices.SortFunc(rds, func(a, b rd) int {
		if a.rad < b.rad {
			return -1
		} else if a.rad > b.rad {
			return 1
		}

		if a.n < b.n {
			return -1
		} else if a.n > b.n {
			return 1
		}

		return 0
	})

	result = strconv.Itoa(rds[limit/10].n)
	return
}

type rd struct {
	n, rad int
}

func rad(m map[int]int) int {
	ret := 1
	for k := range m {
		ret *= k
	}

	return ret
}
