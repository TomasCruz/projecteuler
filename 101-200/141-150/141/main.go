package main

import (
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 141; Square Progressive Numbers
A positive integer, n, is divided by d and the quotient and remainder are q and r respectively.
In addition d, q, and r are consecutive positive integer terms in a geometric sequence, but not necessarily in that order.

For example, 58 divided by 6 has quotient 9 and remainder 4. It can also be seen that 4, 6, 9
are consecutive terms in a geometric sequence (common ratio 3/2).

We will call such numbers, n, progressive.
Some progressive numbers, such as 9 and 10404 = 102^2, happen to also be perfect squares.
The sum of all progressive perfect squares below one hundred thousand is 124657.

Find the sum of all progressive perfect squares below one trillion (10^12).
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
		limit = 12
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	powerThird := uint64(1)
	for i := 0; i < limit/3; i++ {
		powerThird *= 10
	}

	powerHalf := powerThird
	for i := limit / 3; i < limit/2; i++ {
		powerHalf *= 10
	}

	power := powerHalf
	for i := limit / 2; i < limit; i++ {
		power *= 10
	}

	sq := make([]uint64, powerHalf)
	for i := uint64(1); i < powerHalf; i++ {
		sq[i] = i * i
	}

	cubes := make([]uint64, powerThird)
	for i := uint64(1); i < powerThird; i++ {
		cubes[i] = sq[i] * i
	}

	progressivePerfectSquares := []pps{}

	// r = c*b^2, q = a*c*b, d = c*a^2 and n = a^3*b*c^2 + c*b^2
	for a := uint64(2); a < powerThird; a++ {
		for b := uint64(1); b < a; b++ {
			if cubes[a]*b > power {
				break
			}

			if !coprime(a, b) {
				continue
			}

			r := sq[b]
			q := a * b
			d := sq[a]
			for c := uint64(1); ; c++ {
				nCandidate := cubes[a]*b*sq[c] + c*sq[b]
				if nCandidate > power {
					break
				}

				if _, found := slices.BinarySearch(sq, nCandidate); found {
					progressivePerfectSquares = append(progressivePerfectSquares, pps{
						n: nCandidate,
						d: d * c,
						q: q * c,
						r: r * c,
					})
				}
			}
		}
	}

	// slices.SortFunc(progressivePerfectSquares, func(a, b pps) int {
	// 	if a.n < b.n {
	// 		return -1
	// 	} else if a.n > b.n {
	// 		return 1
	// 	}

	// 	return 0
	// })

	sum := uint64(0)

	for _, pps := range progressivePerfectSquares {
		// fmt.Printf("%d / %d = %d (%d)\n", pps.n, pps.d, pps.q, pps.r)
		sum += pps.n
	}

	result = strconv.FormatUint(sum, 10)
	return
}

type pps struct {
	n, d, q, r uint64
}

// modified binary GCD algorithm (Stein's algorithm)
func coprime(a, b uint64) bool {
	if b > a {
		return coprime(b, a)
	}

	// gcd(u,0)=u
	if b == 0 {
		return a == 1
	}

	oddA := a&1 == 1
	oddB := b&1 == 1

	// gcd(2u,2v)=2 * gcd(u,v)
	if !oddA && !oddB {
		return false
	}

	// gcd(u,2v)=gcd(u,v)
	if !oddA {
		return coprime(a>>1, b)
	}

	if !oddB {
		return coprime(a, b>>1)
	}

	// gcd(u,v)=gcd(u,v-u)
	return coprime(a-b, b)
}

/*
	n:d = q(r), n = d*q + r, n > d > r
	n = x^2

	None of d, q or r can't be 0 as there wouldn't be geometric progression.
	r can't be greater than q, (q < r < d), then r = r/q*q, d = r/q*r = r^2/q.
	Then n = r^2/q * q + r = r^2 + r, but then n can't be perfect square.

	So, n > d >= q > r or n > q >= d > r but d*q = q*d i.e. it can be assumed without loss of generality n > d > q > r
	q = q/r * r, d = q/r * q = q^2/r
	n = q^2/r * q + r = q^3/r + r. Considering q/r (geometric progression coefficient greater than 1) is a rational number,
	it can without loss of generality be expessed as a/b, a and b being relatively prime.

	q = r * a/b, d = r * a^2/b^2, and for d to be a natural number, r has to be divisible by b^2 so r = c*b^2.
	d = r * a^2/b^2 = c * b^2 * a^2/b^2 = c*a^2, q = c*b^2 * a/b = a*c*b, r = c*b^2.
	n = c*a^2 * a*c*b + c*b^2 = a^3*b*c^2 + c*b^2.

	Therefore, r = c*b^2, q = a*c*b, d = c*a^2 and n = a^3*b*c^2 + c*b^2.

	a >= 2, if a was 1 then d=c and r=c*b^2, contradicting d > r. 10^12 = n = a^3*b*c^2 + c*b^2 > a^3 => a < 10^4
	b >=1 and b < a (as a/b > 1)

	2 <= a < 10^4, 1 <= b < a
*/
