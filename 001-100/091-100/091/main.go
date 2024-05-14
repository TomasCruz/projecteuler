package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 91; Right Triangles with Integer Coordinates
The points P(x1, y1) and Q(x2, y2) are plotted at integer co-ordinates and are joined to the origin, O(0,0), to form triangle OPQ.

There are exactly fourteen triangles containing a right angle that can be formed when each co-ordinate lies between 0 and 2 inclusive;
that is, 0 <= x1, y1, x2, y2 <= 2.

Given that 0 <= x1, y1, x2, y2 <= 50, how many right triangles can be formed?
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
		limit = 50
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	pCount := 0
	for xq := 0; xq <= limit; xq++ {
		for yq := 0; yq <= limit; yq++ {
			if xq == 0 && yq == 0 {
				continue
			}

			// (x - xq/2)^2 + (y - yq/2)^2 = 1/4 * (xq^2 + yq^2)
			// multiply equation with 4 to use integers
			// (2*x - xq)^2 + (2*y - yq)^2 = xq^2 + yq^2
			r := xq*xq + yq*yq
			for xp := 0; xp <= limit; xp++ {
				for yp := 0; yp <= limit; yp++ {
					xDiff := 2*xp - xq
					yDiff := 2*yp - yq
					leftSide := xDiff*xDiff + yDiff*yDiff
					if leftSide == r {
						pCount++
					}
				}
			}

			pCount -= 2
		}
	}

	resInt := pCount + limit*limit
	result = strconv.Itoa(resInt)

	return
}
