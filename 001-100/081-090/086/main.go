package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 86; Cuboid Route
A spider, S, sits in one corner of a cuboid room, measuring 6 by 5 by 3, and a fly, F, sits in the opposite corner.
By travelling on the surfaces of the room the shortest "straight line" distance from S to F is 10 and the path is shown on the diagram.

However, there are up to three "shortest" path candidates for any given cuboid and the shortest route doesn't always have integer length.

It can be shown that there are exactly 2060 distinct cuboids, ignoring rotations, with integer dimensions, up to a maximum size of M by M by M,
for which the shortest route has integer length when M = 100. This is the least value of M for which the number of solutions
first exceeds two thousand; the number of solutions when M = 99 is 1975.

Find the least value of M such that the number of solutions first exceeds one million.
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

	solutionCount := []int{0, 0}
	epsilon := 0.000001
	m := 2

	for ; ; m++ {
		count := 0
		a := m
		for b := 1; b <= a; b++ {
			for c := 1; c <= b; c++ {
				// all three "shortest" paths will be square root of a*a + b*b + c*c + one of 2*a*b, 2*a*c or 2*b*c
				// b and c being numbers not greater than a will produce smallest product, so shortest path is 2*b*c + a*a + b*b + c*c
				minRoute := 2*b*c + a*a + b*b + c*c
				sqRoot := math.Sqrt(float64(minRoute))
				if math.Abs(sqRoot-math.Round(sqRoot)) < epsilon {
					count++
				}
			}
		}

		newCount := solutionCount[len(solutionCount)-1] + count
		solutionCount = append(solutionCount, newCount)
		if newCount > limit {
			break
		}
	}

	result = strconv.Itoa(m)
	return
}
