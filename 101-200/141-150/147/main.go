package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 147; Rectangles in Cross-hatched Grids
In a 3 * 2 cross-hatched grid, a total of 37 different rectangles could be situated within that grid as indicated in the sketch.

There are 5 grids smaller than 3 * 2, vertical and horizontal dimensions being important,
i.e. 1 * 1, 2 * 1, 3 * 1, 1 * 2 and 2 * 2. If each of them is cross-hatched, the following number of different rectangles
could be situated within those smaller grids:

1 * 1	1
2 * 1	4
3 * 1	8
1 * 2	4
2 * 2	18

Adding those to the 37 of the 3 * 2 grid, a total of 72 different rectangles could be situated within 3 * 2 and smaller grids.

How many different rectangles could be situated within 47 * 43 and smaller grids?
*/

func main() {
	var limitVer, limitHor int

	if len(os.Args) > 1 {
		limit64Ver, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limitVer = int(limit64Ver)

		limit64Hor, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limitHor = int(limit64Hor)
	} else {
		limitVer = 47
		limitHor = 43
	}

	projecteuler.Timed(calc, limitVer, limitHor)
}

func calc(args ...interface{}) (result string, err error) {
	limitVer := args[0].(int)
	limitHor := args[1].(int)

	rectCount := make([][]int, limitVer+1)
	for i := 0; i <= int(limitVer); i++ {
		rectCount[i] = make([]int, limitHor+1)
	}

	for v := 1; v <= limitVer; v++ {
		for h := 1; h <= limitHor; h++ {
			currCount := 0
			maxVH := v + h - 2
			for ad := 1; ad <= maxVH; ad++ {
				for md := 1; md <= maxVH; md++ {
					currCount += diagCount(v, h, ad, md)
				}
			}

			for i := 0; i < v; i++ {
				for j := 0; j < h; j++ {
					currCount += (v - i) * (h - j)
				}
			}

			rectCount[v][h] = currCount
		}
	}

	sum := int64(0)
	for i := 1; i <= limitVer; i++ {
		for j := 1; j <= limitHor; j++ {
			sum += int64(rectCount[i][j])
		}
	}

	result = strconv.FormatInt(sum, 10)
	return
}

type point [2]int
type rect [4]point

func diagCount(v, h, ad, md int) int {
	/*
		3, 4, 2, 3				3, 4, 3, 2				3, 4, 1, 4

		*   1   *   *   *		*   *   1   *   *		*   1   *   *   *
		  *   *   *   *			  *   *   *   *			  1   *   *   *
		1   *   *   *   *		*   *   *   1   *		*   *   *   *   *
		  *   *   1   *			  1   *   *   *			  *   *   *   *
		*   *   *   *   *		*   *   *   *   *		*   *   *   1   *
		  *   1   *   *			  *   1   *   *			  *   *   1   *
		*   *   *   *   *		*   *   *   *   *		*   *   *   *   *
	*/

	lu := point{0, ad}
	if ad%2 == 1 {
		lu[1]++
	}

	ld := point{ad, 0}
	if ad%2 == 1 {
		ld[1]++
	}

	rd := point{ld[0] + md, ld[1] + md}
	ru := point{lu[0] + md, lu[1] + md}

	r := rect{lu, ld, rd, ru}
	r1 := r.move()

	return diagCountX(v, h, r) + diagCountX(v, h, r1)
}

func (r rect) move() rect {
	ret := rect{r[0], r[1], r[2], r[3]}

	if r[1][1] == 1 {
		ret[0][1]--
		ret[1][1]--
		ret[2][1]--
		ret[3][1]--
	} else {
		ret[0][1]++
		ret[1][1]++
		ret[2][1]++
		ret[3][1]++
	}

	ret[0][0]++
	ret[1][0]++
	ret[2][0]++
	ret[3][0]++

	return ret
}

func diagCountX(v, h int, r rect) int {
	maxH := 2 * h
	if r[3][0]%2 == 1 {
		maxH--
	}
	if r[3][1] > maxH {
		return 0
	}
	hMul := (maxH-r[3][1])/2 + 1

	maxV := 2 * v
	if r[2][0]%2 == 1 {
		maxV--
	}
	if r[2][0] > maxV {
		return 0
	}
	vMul := (maxV-r[2][0])/2 + 1

	return hMul * vMul
}
