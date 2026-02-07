package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 134; Prime Pair Connection
Consider the consecutive primes p_1 = 19 and p_2 = 23. It can be verified that 1219 is the smallest number such that
the last digits are formed by p_1 whilst also being divisible by p_2.

In fact, with the exception of p_1 = 3 and p_2 = 5, for every pair of consecutive primes, p_2 > p_1, there exist
values of n for which the last digits are formed by p_1 and n is divisible by p_2. Let S be the smallest of these values of n.

Find sum(S) for every pair of consecutive primes with 5 <= p_1 <= 1000000.
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
		limit = 10
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := uint64(args[0].(int))

	primes := projecteuler.PrimesEratosthenes(limit+100, nil)

	sum := uint64(0)
	power := uint64(10)
	for i := 2; i < len(primes) && primes[i] <= limit; i++ {
		if primes[i] > power {
			power *= 10
		}

		rPower := power % primes[i+1]

		// min x for x*rPower + primes[i] = 0 (mod primes[i+1])
		// x*rPower = primes[i+1] - primes[i] (mod primes[i+1])
		// a*x = b (mod n) [ a = rPower, b = primes[i+1] - primes[i], n = primes[i+1] ]
		// x = a^(-1)*b (mod n)
		inverse := extendedEuclidean(int(rPower), int(primes[i+1]))
		x := (uint64(inverse) * (primes[i+1] - primes[i])) % primes[i+1]
		s := power*x + primes[i]
		sum += s
	}

	result = strconv.FormatUint(sum, 10)
	return
}

// extended Euclidean algorithm
func extendedEuclidean(a, n int) int {
	r, nr := n, a
	t, nt := 0, 1
	for nr != 0 {
		q := r / nr
		r, nr = nr, r-q*nr
		t, nt = nt, t-q*nt
	}

	if t < 0 {
		t += n
	}

	return t
}

/*
	n for which the last digits are formed by p_1 and n is divisible by p_2

	S(19, 23) = 1219 => y = min(p2 | [x*10^numDigs(p1) + p1]) = 1219
	10^numDigs(p1) = y (mod p2), x*10^numDigs(p1) = x*y (mod p2), x*y + p1 = 0 (mod p2)

	x*100 + 19 = 0 (mod 23), 100 = 8 (mod 23)
	8x + 19 = 0 (mod 23)
	for x 1:4, 2:12, 3:20, 4: 5, 5:13, 6:21, 7:6, 8:14, 9:22, 10:7, 11:15, 12:0, so 1219

	8x = -19(23), 8x = 4(23)
*/
