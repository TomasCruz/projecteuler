package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 149; Maximum-sum Subsequence
Looking at the table below, it is easy to verify that the maximum possible sum of adjacent numbers in any direction
(horizontal, vertical, diagonal or anti-diagonal) is 16 (= 8 + 7 + 1).

-2    5    3    2
 9   -6    5    1
 3    2    7    3
-1    8   -4    8

Now, let us repeat the search, but on a much larger scale:

First, generate four million pseudo-random numbers using a specific form of what is known as a "Lagged Fibonacci Generator":

For 1 <= k <= 55, s(k) = [100003 - 200003*k + 300007*k^3] % 1000000 - 500000.
For 56 <= k <= 4000000, s(k) = [s(k-24) + s(k-55) + 1000000] % 1000000 - 500000.

Thus, s(10) = -393027 and s(100) = 86613.

The terms of s are then arranged in a 2000 * 2000 table, using the first 2000 numbers to fill the first row (sequentially),
the next 2000 numbers to fill the second row, and so on.

Finally, find the greatest sum of (any number of) adjacent entries in any direction (horizontal, vertical, diagonal or anti-diagonal).
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
		limit = 2000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	m := generateMatrix(limit)

	maxes := make([]int64, 4)
	maxes[0] = generateHorizontal(limit, m)
	maxes[1] = generateVertical(limit, m)
	maxes[2] = generateDiagonal(limit, m)
	maxes[3] = generateAntiDiagonal(limit, m)

	max := int64(math.MinInt64)
	for i := range 4 {
		if maxes[i] > max {
			max = maxes[i]
		}
	}

	result = strconv.FormatInt(max, 10)
	return
}

func generateMatrix(limit int) [][]int {
	if limit == 4 {
		return [][]int{
			{0, 0, 0, 0, 0},
			{0, -2, 5, 3, 2},
			{0, 9, -6, 5, 1},
			{0, 3, 2, 7, 3},
			{0, -1, 8, -4, 8},
		}
	}

	s := make([]int, limit*limit+1)

	for k := 1; k < 56; k++ {
		k3 := k * k * k
		s[k] = (100003-200003*k+300007*k3)%1000000 - 500000
	}

	for k := 56; k < limit*limit+1; k++ {
		s[k] = (s[k-24]+s[k-55]+1000000)%1000000 - 500000
	}

	ret := make([][]int, limit+1)
	for row := range limit + 1 {
		ret[row] = make([]int, limit+1)
	}

	for row := 1; row <= limit; row++ {
		for col := range limit + 1 {
			k := (row-1)*limit + col
			ret[row][col] = s[k]
		}
	}

	return ret
}

func generateHorizontal(limit int, m [][]int) int64 {
	a := make([][]int64, limit+1)
	for row := range limit + 1 {
		a[row] = make([]int64, limit+1)
	}

	for i := 1; i <= limit; i++ {
		a[i][1] = int64(m[i][1])
	}

	maxes := make([]int64, limit+1)

	for row := 1; row <= limit; row++ {
		mmm := int64(math.MinInt64)

		for col := 1; col < limit; col++ {
			mm := int64(m[row][col+1])
			if a[row][col] > 0 {
				mm += a[row][col]
			}
			a[row][col+1] = mm

			if a[row][col+1] > mmm {
				mmm = a[row][col+1]
			}
		}

		maxes[row] = mmm
	}

	ret := int64(maxes[1])
	for i := 2; i <= limit; i++ {
		if maxes[i] > ret {
			ret = maxes[i]
		}
	}

	return ret
}

func generateVertical(limit int, m [][]int) int64 {
	a := make([][]int64, limit+1)
	for row := range limit + 1 {
		a[row] = make([]int64, limit+1)
	}

	for i := 1; i <= limit; i++ {
		a[1][i] = int64(m[1][i])
	}

	maxes := make([]int64, limit+1)

	for col := 1; col <= limit; col++ {
		mmm := int64(math.MinInt64)

		for row := 1; row < limit; row++ {
			mm := int64(m[row+1][col])
			if a[row][col] > 0 {
				mm += a[row][col]
			}
			a[row+1][col] = mm

			if a[row+1][col] > mmm {
				mmm = a[row+1][col]
			}
		}

		maxes[col] = mmm
	}

	ret := int64(maxes[1])
	for i := 2; i <= limit; i++ {
		if maxes[i] > ret {
			ret = maxes[i]
		}
	}

	return ret
}

func generateDiagonal(limit int, m [][]int) int64 {
	a := make([][]int64, limit+1)
	for row := range limit + 1 {
		a[row] = make([]int64, limit+1)
	}

	for i := 1; i <= limit; i++ {
		a[i][1] = int64(m[i][1])
		a[1][i] = int64(m[1][i])
	}

	maxes := make([]int64, 2*limit)
	maxes[2*limit-1] = int64(m[1][limit])

	for row := 1; row < limit; row++ {
		mmm := int64(math.MinInt64)

		for col := 1; col < limit; col++ {
			mm := int64(m[row+1][col+1])
			if a[row][col] > 0 {
				mm += a[row][col]
			}
			a[row+1][col+1] = mm

			if a[row+1][col+1] > mmm {
				mmm = a[row+1][col+1]
			}
		}

		maxes[row] = mmm
	}

	for row := limit + 1; row < 2*limit-1; row++ {
		mmm := int64(math.MinInt64)

		for col := row - limit + 1; col < limit; col++ {
			mm := int64(m[row-limit+1][col+1])
			if a[row-limit][col] > 0 {
				mm += a[row-limit][col]
			}
			a[row-limit+1][col+1] = mm

			if a[row-limit+1][col+1] > mmm {
				mmm = a[row-limit+1][col+1]
			}
		}

		maxes[row-limit] = mmm
	}

	ret := int64(m[limit][1])
	for i := 2; i < 2*limit; i++ {
		if maxes[i] > ret {
			ret = maxes[i]
		}
	}

	return ret
}

func generateAntiDiagonal(limit int, m [][]int) int64 {
	a := make([][]int64, limit+1)
	for row := range limit + 1 {
		a[row] = make([]int64, limit+1)
	}

	for i := 1; i < limit; i++ {
		a[i][1] = int64(m[i][1])
		a[limit][i] = int64(m[limit][i])
	}

	maxes := make([]int64, 2*limit)
	maxes[2*limit-1] = int64(m[limit][limit])

	for row := 2; row <= limit; row++ {
		mmm := int64(math.MinInt64)

		for col := 1; col < row; col++ {
			mm := int64(m[row-col][col+1])
			if a[row-col+1][col] > 0 {
				mm += a[row-col+1][col]
			}
			a[row-col][col+1] = mm

			if a[row-col][col+1] > mmm {
				mmm = a[row-col][col+1]
			}
		}

		maxes[row] = mmm
	}

	for row := limit + 1; row < 2*limit-1; row++ {
		mmm := int64(math.MinInt64)

		for col := row - limit + 1; col < limit; col++ {
			mm := int64(m[row-col][col+1])
			if a[row-col+1][col] > 0 {
				mm += a[row-col+1][col]
			}
			a[row-col][col+1] = mm

			if a[row-col][col+1] > mmm {
				mmm = a[row-col][col+1]
			}
		}

		maxes[row] = mmm
	}

	ret := int64(m[1][1])
	for i := 2; i < 2*limit; i++ {
		if maxes[i] > ret {
			ret = maxes[i]
		}
	}

	return ret
}
