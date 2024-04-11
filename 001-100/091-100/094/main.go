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
	result = strconv.Itoa(int(sum))

	return
}

func perimeterSumAET(limit int64) int64 {
	tree := projecteuler.NewFibonacciBoxTernaryTree(1, 1, 2, 3)
	tree.Generate(int(limit / 2))
	triplets := tree.TripletSlice()

	perimeterSum := int64(0)
	for _, t := range triplets {
		lesserArm := t.A
		if t.B < t.A {
			lesserArm = t.B
		}

		if 2*lesserArm+1 == t.C || 2*lesserArm-1 == t.C {
			perimeterSum += int64(2 * (lesserArm + t.C))
		}
	}

	return perimeterSum
}
