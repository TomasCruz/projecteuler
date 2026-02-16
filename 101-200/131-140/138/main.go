package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 138; Special Isosceles Triangles
Consider the isosceles triangle with base length, b = 16, and legs, L = 17.

By using the Pythagorean theorem it can be seen that the height of the triangle, h = sqrt(17^2 - 8^2) = 15, which is one less than the base length.

With b = 272 and L = 305, we get h = 273, which is one more than the base length, and this is the second smallest isosceles triangle
with the property that h = b +- 1.

Find sum(L) for the twelve smallest isosceles triangles for which h = b +- 1 and b, L are positive integers.
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

	x := uint64(2)
	l := uint64(1)

	sum := uint64(0)
	for i := 0; i < limit; {
		x, l = 9*x+20*l, 4*x+9*l

		switch x % 5 {
		case 2:
			sum += l
			i++
		case 3:
			sum += l
			i++
		}
	}

	result = strconv.FormatUint(sum, 10)
	return
}

/*
	examples:
	b = 16, L = 17, h = 15
	b = 272, L = 305, h = 273

	L^2 = h^2 + (b/2)^2, h = b +- 1
	L^2 = b^2 +- 2*b + 1 + b^2/4
	4L^2 = 4b^2 +- 8*b + 4 + b^2
	4L^2 = 5b^2 +- 8*b + 4					(1)

	5b^2 +- 8*b + 4 - 4L^2 = 0
	b1,2 = (8 +- sqrt(64 - 20*(4 - 4*L^2))) / 10
	b1,2 = (8 +- sqrt(80*L^2 - 16))) / 10
	b1,2 = (4 +- 2*sqrt(5*L^2 - 1))) / 5
	b3,4 = (-4 +- 2*sqrt(5*L^2 - 1))) / 5

	Since b has to be positive,
	b1,2 = (2*sqrt(5*L^2 - 1)) +- 4) / 5	(2)

	x := sqrt(5*L^2 - 1))
	5*L^2 - 1 = x^2							(3)

	b1,2 = (2*x +- 4) / 5
	x = (5*b +- 4)/2
	x = 5/2*b +- 2							(4)

	Since b is an even number and 5*b +- 4 > 0, x is a natural number

	x^2 mod 5 = 0,1,4,4,1 possibilities corresponding to x mod 5 = 0,1,2,3,4
	From (3), x mod 5 = 2,3 so x^2 mod 5 = 4
	(x^2 + 1) mod 5 = 0

	b = (2x +- 4) / 5
	2x mod 5 = 4,1
	(2x + 4) mod 5 = 3,0
	(2x - 4) mod 5 = 0,2

	x mod 5 = 2, b = (2x - 4) / 5, h = b + 1, l = sqrt((x^2 + 1) / 5)
	x mod 5 = 3, b = (2x + 4) / 5, h = b - 1, l = sqrt((x^2 + 1) / 5)

	x = 38, b = 16, h = 15, l = 17
	x = 682, b = 272, h = 273, l = 305

	Multiplying (1) by 5/4,
	4L^2 = 5b^2 +- 8*b + 4 / * (5/4)
	5L^2 = 25/4*b^2 +- 10*b + 5 = (5/2*b +- 2)^2 + 1
	5*L^2 - 1 = (5/2*b +- 2)^2
	(5/2*b +- 2)^2 - 5*L^2 = -1				(5)

	Replacing (4) in (5),
	x^2 - 5*L^2 = -1						(6)		Pell's negative equation

	2^2 - 5*1 = -1, trivial solution s(1) is (2, 1)

	All positive solutions are odd powers of (2 + sqrt(5))
	s(1) = (2, 1)
	s(n) = (x(n), y(n)) is such that x(n) + sqrt(5)*y(n) = (2 + sqrt(5))^(2n-1)
	x(n) + sqrt(5)*y(n) = (2 + sqrt(5))^(2n-1) = (2 + sqrt(5))^2 * (2 + sqrt(5))^(2n-3)
	x(n) + sqrt(5)*y(n) = (9 + 4*sqrt(5)) * (x(n-1) + sqrt(5)*y(n-1))
	x(n) + sqrt(5)*y(n) = 9x(n-1) + 20y(n-1) + sqrt(5)*(4x(n-1) + 9y(n-1))

	x(n) = 9x(n-1) + 20y(n-1)
	y(n) = 4x(n-1) + 9y(n-1)
*/
