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
Problem 67; Maximum path sum II

By starting at the top of the triangle below and moving to adjacent numbers on the row below,
the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'),
a 15K text file containing a triangle with one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible to try every route to solve this problem,
as there are 2^99 altogether! If you could check one trillion (10^12) routes every second it would take over
twenty billion years to check them all. There is an efficient algorithm to solve it. ;o)
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
		limit = 100
	}

	projecteuler.TimedStr(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	var textNumbers []string
	if textNumbers, err = projecteuler.FileToStrings("p067_triangle.txt"); err != nil {
		fmt.Println(err)
		return
	}

	var m [][]int
	if m, err = buildMatrix(limit, textNumbers); err != nil {
		return
	}

	res := buildAddedMatrix(limit, m)
	//projecteuler.PrintMatrix(result)

	result = strconv.Itoa(res[0][0])
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
