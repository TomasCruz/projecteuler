package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 85; Counting Rectangles
By counting carefully it can be seen that a rectangular grid measuring 3 by 2 contains eighteen rectangles.
Although there exists no rectangular grid that contains exactly two million rectangles, find the area of the grid with the nearest solution.
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
		limit = 2000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	/*
		For rect widt w, inner rects can be size 1..w
		rc := sum i 1..w [(1 + w - i)] * sum j 1..h [1 + h - j]
		sum j 1..h [1 + h - j] == h*(1+h) - sum j 1..h [j] == h*(1+h)/2

		rc = w*(1+w)/2 * h*(1+h)/2
		e.g. w=3, h=2: 6*3 = 18
	*/

	limit := args[0].(int)

	nearestDiff := 2 * limit
	nearestW := 0
	nearestH := 0

	for w := 1; ; w++ {
		var h int
		for h = 1; h <= w; h++ {
			prod := (w * (1 + w) * h * (1 + h)) / 4
			d := limit - prod
			if d < 0 {
				d = -d
			}

			if d < nearestDiff {
				nearestDiff = d

				nearestW = w
				nearestH = h
			}

			if prod > limit {
				break
			}
		}

		if h == 1 {
			break
		}
	}

	result = strconv.Itoa(nearestW * nearestH)
	return
}
