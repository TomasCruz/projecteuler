package main

import (
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 131; Prime Cube Partnership
There are some prime values, p, for which there exists a positive integer, n, such that the expression n^3 + n^2*p is a perfect cube.

For example, when p = 19, 8^3 + 8^2 * 19 = 12^3.

What is perhaps most surprising is that for each prime with this property the value of n is unique, and there are only four such primes below one-hundred.

How many primes below one million have this remarkable property?
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
		limit = 6
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	power := 1
	for i := 0; i < limit; i++ {
		power *= 10
	}

	primes := projecteuler.Primes(power, nil)

	res := []int{}

	for a := 1; ; a++ {
		pCandidate := 3*a*(a+1) + 1
		if pCandidate > power {
			break
		}

		if _, found := slices.BinarySearch(primes, pCandidate); found {
			res = append(res, pCandidate)
		}
	}

	result = strconv.Itoa(len(res))
	return
}

/*
	n^3 + n^2*p = x^3			(1)
	n^2*(n + p) = x^3			(2)

	If p|n, then n = k*p for some natural k.
	Replacing in (2), k^2*p^2*(k*p + p) = k^2*p^3*(k + 1) = x^3 => p^3 | x^3 => p|x
	Since x = l*p (for some l natural), k^2*p^3*(k + 1) = l^3*p^3
	k^2*(k + 1) = l^3. So l > k, l = k + d, for some natural d
	k^3 + k^2 = (k + d)^3 = k^3 + 3*k^2*d + 3*k*d^2 + d^3
	k^2*(3d - 1) + 3*k*d^2 + d^3 = 0
	as d >= 1, 3d-1 >= 2, 3*d^2 >= 3, d^3 >= 1
	Quadratic equation with positive coefficients doesn't have positive solutions, so n is not divisible by p.
	n/p can't be reduced, and (n + p)/n also can't be reduced.

	From (1), n^3*(1 + p/n) = x^3. Applying cubic root,
	n * [(n+p)/n]^(1/3) = x. (n+p)/n has to be a perfect cube, so both n+p and n have to be perfect cubes.

	n = a^3, n+p = b^3 => b > a, both a and b being natural numbers
	Replacing in (2), (a^3)^2 * b^3 = x^3 => x = a^2*b
	(a^3)^2 * (a^3 + p) = a^6*b^3
	Dividing by a^6, a^3 + p = b^3
	p = b^3 - a^3				(4)

	b = a + d for some natural d, p = (a+d)^3 - a^3 = 3a^2*d + 3a*d^2 + d^3 = d*(3a^2 + 3a*d + d^2)
	For natural a and d, 3a^2 + 3a*d + d^2 will always be greater than 1, therefore d=1.
	Replacing in (4), p = (a+1)^3 - a^3 = a^3 + 3a^2 + 3a + 1 - a^3 = 3a^2 + 3a + 1

	p = 3a*(a + 1) + 1			(5)

	n = a^3, n+p = (a+1)^3, x = a^2*(a+1)
	(a^3)^2 * ((a^3)^3 + p) = (a^2*(a+1))^3
*/
