package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 83; Path Sum: Four Ways

NOTE: This problem is a significantly more challenging version of Problem 81.

In the 5 by 5 matrix below, the minimal path sum from the top left to the bottom right, by moving left, right, up, and down, is indicated with asterisk
and is equal to 2297.

	*131	673		*234	*103	*18
	*201	*96		*342	965		*150
	630		803		746		*422	*111
	537		699		497		*121	956
	805		732		524		*37		*331

Find the minimal path sum from the top left to the bottom right by moving left, right, up, and down in
matrix.txt (right click and "Save Link/Target As..."), a 31K text file containing an 80 by 80 matrix.
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
		fileName = "0083_matrix.txt"
	}

	var rowStrings []string
	if rowStrings, err = projecteuler.FileToStrings(fileName); err != nil {
		return
	}

	input := makeMatrix(limit)
	for i := 0; i < limit; i++ {
		readToRow(limit, rowStrings[i], input[i])
	}

	// min path matrix
	minPath := newMinPathMatrix(limit, input)
	minPath.buildMinPathMatrix()

	result = strconv.FormatInt(minPath.m[limit-1][limit-1], 10)
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
