package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 94; Almost Equilateral Triangles
It is easily proved that no equilateral triangle exists with integral length sides and integral area. However, the almost equilateral triangle
5-5-6 has an area of 12 square units.

We shall define an almost equilateral triangle to be a triangle for which two sides are equal and the third differs by no more than one unit.
Find the sum of the perimeters of all almost equilateral triangles with integral side lengths and area and whose perimeters
do not exceed one billion (1,000,000,000).
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(...interface{}) (result string, err error) {
	limit := int64(1000000000)

	sum := perimeterSumAET(limit)
	result = strconv.FormatInt(sum, 10)

	return
}

/*
Initial brutish approach too slow. From user "Ende" from projecteuler.net:

I used Pell equations, my new favorite toy. If they could make steak, I'd use them to do that, too.

Code is long and perhaps not very clear, but this is the theory behind it:

Consider half of an almost equilateral triangle (AET), obtaining a right triangle with sides a, b, c (c being a side from the original
almost equilateral triangle, and b being the height on the symmetric axis).

Gamma illustrated it like this earlier in the thread:
In Ascii-art:
C/|B
/_|
A
Extend this to:
C/|\ C
/_|_\
A A

For a, b, c to actually be half of an AET, we need a^2 + b^2 = c^2 to hold, as well as c = 2a +- 1.

Starting from

a^2 + b^2 = 2a +- 1,

we can get to

u^2 - 3b^2 = 1, where u = 3(a +- 2/3), which is a lovely Pell equation. :)


You would think that c = 2a+2 would have to be considered as well, which could then be scaled down to an almost equilateral triangle,
but it turns out that for that, there are no solutions to the respective Pell equation.
*/

func perimeterSumAET(limit int64) int64 {
	// a^2 + b^2 = (2a +- 1)^2
	// a^2 + b^2 = 4a^2 +- 4a + 1
	// 3a^2 +- 4a - b^2 = -1
	// 3[a^2 +- 4/3a] - b^2 = -1		//// (*3)
	// [9*(a^2 +- 4/3a)] - 3b^2 = -3	//// (+4)
	// [9*(a^2 +- 4/3a + 4/9)] - 3b^2 = 1
	// [3*(a +- 2/3)]^2 - 3b^2 = 1
	// u^2 - 3b^2 = 1, where u = 3a +- 2

	fundamental := projecteuler.Chakravala(3)

	// for u = 3*a - 2, arm is 2*a - 1, perimeter is 6*a - 2, and for u = 3*a + 2, arm is 2*a + 1, perimeter is 6*a + 2
	// So, if u mod 3 = 1, p = 6a - 2 and if u mod 3 = 2, p = 6a + 2
	// According to Θαλῆς, a + (a +- 1) >  b, otherwise numerical solution of Pell's equation doesn't correspond to a triangle

	sum := int64(0)
	var nT, prev projecteuler.ChakravalaTriplet
	prev = fundamental
	for {
		nT = projecteuler.Samasa(3, fundamental.X, prev)
		prev = nT

		u := nT.X.Int64()
		b := nT.Y.Int64()

		if u%3 == 1 {
			a := (u + 2) / 3
			if 3*a-1 > b {
				p := 6*a - 2
				if p > limit {
					break
				}
				// fmt.Println("minus", 2*a, 2*a-1, b, p)
				sum += p
			}
		} else if u%3 == 2 {
			a := (u - 2) / 3
			if 3*a+1 > b {
				p := 6*a + 2
				if p > limit {
					break
				}
				// fmt.Println("plus", 2*a, 2*a+1, b, p)
				sum += p
			}
		}
	}

	return sum
}
