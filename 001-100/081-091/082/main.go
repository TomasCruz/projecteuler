package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 82; Path Sum: Three Ways

NOTE: This problem is a more challenging version of Problem 81.

The minimal path sum in the 5 by 5 matrix below, by starting in any cell in the left column and finishing in any cell in the right column,
and only moving up, down, and right, is indicated with asterisk; the sum is equal to 994.

	131		673		*234	*103	*18
	*201	*96		*342	965		150
	630		803		746		422		111
	537		699		497		121		956
	805		732		524		37		331

Find the minimal path sum from the left column to the right column
in matrix.txt (right click and "Save Link/Target As..."), a 31K text file containing an 80 by 80 matrix.
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
		limit = 80
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	var fileName string
	if limit == 5 {
		fileName = "example.txt"
	} else {
		fileName = "0082_matrix.txt"
	}

	var rowStrings []string
	if rowStrings, err = projecteuler.FileToStrings(fileName); err != nil {
		return
	}

	input := makeMatrix(limit)
	for i := 0; i < limit; i++ {
		readToRow(limit, rowStrings[i], input[i])
	}

	minResult := int64(math.MaxInt64)
	for row := 0; row < limit; row++ {
		// min path matrix
		minPath := newMinPathMatrix(limit, input)
		minPath.buildMinPathMatrix(row, 0)

		currMinPath := int64(math.MaxInt64)
		for i := 0; i < limit; i++ {
			if minPath.m[i][limit-1] < currMinPath {
				currMinPath = minPath.m[i][limit-1]
			}
		}

		if currMinPath < minResult {
			minResult = currMinPath
		}
	}

	result = strconv.FormatInt(minResult, 10)
	return
}

func makeMatrix(limit int) [][]int64 {
	m := make([][]int64, limit)
	for i := 0; i < limit; i++ {
		m[i] = make([]int64, limit)
	}

	return m
}

func readToRow(limit int, rowString string, row []int64) (err error) {
	numStrings := strings.Split(rowString, ",")
	for i := 0; i < limit; i++ {
		if row[i], err = strconv.ParseInt(numStrings[i], 10, 64); err != nil {
			return
		}
	}

	return
}
