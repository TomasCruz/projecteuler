package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 18; Maximum path sum I

By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total
from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route.
However, Problem 67, is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force,
and requires a clever method! ;o)
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
		limit = 15
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	var textNumbers []string
	if textNumbers, err = projecteuler.FileToStrings("input.txt"); err != nil {
		fmt.Println(err)
		return
	}

	var m [][]int
	if m, err = buildMatrix(limit, textNumbers); err != nil {
		return
	}

	resultMatrix := buildAddedMatrix(limit, m)
	//projecteuler.PrintMatrix(resultMatrix)

	result = strconv.Itoa(resultMatrix[0][0])
	return
}

func buildMatrix(dim int, textNumbers []string) (m [][]int, err error) {
	m = make([][]int, dim)

	for i := 0; i < dim; i++ {
		m[i] = make([]int, dim)
		numStrings := strings.Split(textNumbers[i], " ")

		var j int
		for j = 0; j < len(numStrings); j++ {
			input := strings.TrimLeftFunc(numStrings[j], func(r rune) bool { return r == '0' })
			if m[i][j], err = strconv.Atoi(input); err != nil {
				fmt.Println(err)
				return
			}
		}

		for ; j < dim; j++ {
			m[i][j] = 0
		}
	}

	return
}

func buildAddedMatrix(dim int, m [][]int) (dp [][]int) {
	dp = make([][]int, dim)

	for i := 0; i < dim; i++ {
		dp[i] = make([]int, dim)
		copy(dp[i], m[i])
	}

	for i := dim - 1; i > 0; i-- {
		for j := 0; j <= i; j++ {
			if j > 0 {
				curr := dp[i][j] + m[i-1][j-1]
				if curr > dp[i-1][j-1] {
					dp[i-1][j-1] = curr
				}
			}

			if j < i {
				dp[i-1][j] = dp[i][j] + m[i-1][j]
			}
		}
	}

	return
}
