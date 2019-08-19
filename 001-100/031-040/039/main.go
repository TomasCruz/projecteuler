package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 39; Integer right triangles

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions
for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
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

type triangle struct {
	a, b, c int
}

func newTriangle(a, b, c int) (t triangle) {
	t = triangle{a: a, b: b, c: c}
	return t
}

func (tr triangle) print() {
	fmt.Println("(", tr.a, ", ", tr.b, ", ", tr.c, ")")
}

func (tr triangle) mul(k int) (nt triangle) {
	nt = triangle{a: tr.a * k, b: tr.b * k, c: tr.c * k}
	return
}

func calc(args ...interface{}) (result string, err error) {
	/*
		a = m^2 - n^2
		b = 2mn
		c = m^2 + n^2
		m > n > 0

		p = a + b + c = 2m(m+n) <= limit (p may be composed by different combinations of m and n!)

		Since p <= 1000 (max limit) => m(m+n) <= 500
		limit/2 >= m(m+n) > n*2n = 2n^2 => n < sqrt(limit/4) => n from [1,15] (for max limit)

		Limiting m: m(m+n) = p/2 means m can be >= n+1, and such that 2*m*(m+n) <= limit
	*/

	limit := args[0].(int)

	pMap := make(map[int]map[triangle]struct{})
	maxN := int(math.Floor(math.Sqrt(float64(limit) / 4)))
	for n := 1; n <= maxN; n++ {
		for m := n + 1; ; m++ {
			// make them coprime to get primitive triples
			cprN, cprM := n, m
			divisor := 2
			for divisor < cprN {
				for cprN%divisor == 0 && cprM%divisor == 0 {
					cprN /= divisor
					cprM /= divisor
				}

				divisor++
			}

			p := 2 * cprM * (cprM + cprN)
			if p > limit {
				break
			}

			a := cprM*cprM - cprN*cprN
			b := 2 * cprM * cprN
			if b < a {
				tmp := a
				a, b = b, tmp
			}
			t := newTriangle(a, b, cprM*cprM+cprN*cprN)

			if _, ok := pMap[p]; !ok {
				pMap[p] = make(map[triangle]struct{})
			}
			pMap[p][t] = struct{}{}
		}
	}

	primes := projecteuler.Primes(limit+1, nil)
	triangleMap := make(map[int]map[triangle]struct{})

	for p := 1; p <= limit; p++ {
		factors, _ := projecteuler.Factorize(p, primes)
		divisors := projecteuler.FindDivisors(factors)

		for i := 1; i < len(divisors); i++ { // disregard divisor 1
			if dividedTriangles, ok := pMap[divisors[i]]; ok {
				if _, ok := triangleMap[p]; !ok {
					triangleMap[p] = make(map[triangle]struct{})
				}

				for dt := range dividedTriangles {
					triangleMap[p][dt.mul(p/divisors[i])] = struct{}{}
				}
			}
		}
	}

	maxP, maxTriangleNumber := 0, 0
	for k, v := range triangleMap {
		triangleNumber := len(v)
		if triangleNumber > maxTriangleNumber {
			maxTriangleNumber = triangleNumber
			maxP = k
		}
	}

	result = strconv.Itoa(maxP)
	return
}
