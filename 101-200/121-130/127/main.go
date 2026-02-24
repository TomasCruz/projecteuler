package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 127; abc-hits
The radical of n, rad(n), is the product of distinct prime factors of n. For example, 504 = 2^3 * 3^2 * 7, so rad(504) = 2 * 3 * 7 = 42.
We shall define the triplet of positive integers (a, b, c) to be an abc-hit if:

gcd(a, b) = gcd(a, c) = gcd(b, c) = 1
a < b
a + b = c
rad(abc) < c

For example, (5, 27, 32) is an abc-hit, because:
gcd(5, 27) = gcd(5, 32) = gcd(27, 32) = 1
5 < 27
5 + 27 = 32
rad(4320) = 30 < 32

It turns out that abc-hits are quite rare and there are only thirty-one abc-hits for c < 1000, with sum(c) = 12523.
Find sum(c) for c < 120000.
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
		limit = 120000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(limit, nil)
	factorSlice := make([]map[int]int, limit)
	radSlice := make([]int, limit)
	for i := 1; i < limit; i++ {
		factorSlice[i], _ = projecteuler.Factorize(i, primes)
		radSlice[i] = rad(factorSlice[i])
	}

	abcHits := []triplet{}

	limitHalf := limit / 2
	for a := 1; a < limitHalf; a++ {
		for b := a + 1; ; b++ {
			c := a + b
			if c >= limit {
				break
			}

			if radSlice[a]*radSlice[b]*radSlice[c] >= c {
				continue
			}

			// insight by user neverforget,
			// Basically, given gcd(a,c)=1 and a+b=c, suppose for contradiction that gcd(a,b)!=1.
			// Then some d>1 must divide both a and b,
			// but then d must divide a+b=c. Similarly gcd(b,c)=1.
			if isDisjunct(factorSlice[a], factorSlice[c]) {
				abcHits = append(abcHits, triplet{a: a, b: b, c: c})
			}
		}
	}

	sum := 0
	for _, abc := range abcHits {
		sum += abc.c
	}

	result = strconv.Itoa(sum)
	return
}

func isDisjunct(l, r map[int]int) bool {
	for k := range l {
		if _, present := r[k]; present {
			return false
		}
	}

	return true
}

type triplet struct {
	a, b, c int
}

func rad(m map[int]int) int {
	ret := 1
	for k := range m {
		ret *= k
	}

	return ret
}
