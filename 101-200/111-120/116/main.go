package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 116; Red, Green or Blue Tiles
A row of five grey square tiles is to have a number of its tiles replaced with coloured oblong tiles chosen
from red (length two), green (length three), or blue (length four).

If red tiles are chosen there are exactly seven ways this can be done.
If green tiles are chosen there are three ways.
And if blue tiles are chosen there are two ways.

Assuming that colours cannot be mixed there are 7 + 3 + 2 = 12 ways of replacing the grey tiles in a row measuring five units in length.

How many different ways can the grey tiles in a row measuring fifty units in length be replaced
if colours cannot be mixed and at least one coloured tile must be used?

NOTE: This is related to Problem 117.
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

	m2 := calcMatrix(2, limit)
	m3 := calcMatrix(3, limit)
	m4 := calcMatrix(4, limit)

	sum := m2 + m3 + m4

	result = strconv.FormatUint(sum, 10)
	return
}

func calcMatrix(m, n int) uint64 {
	maxGroups := n / m

	ret := make([][]uint64, maxGroups+2)
	for row := 0; row < maxGroups+2; row++ {
		ret[row] = make([]uint64, n+1)
	}

	for col := m; col <= n; col++ {
		ret[1][col] = uint64(col - m + 1)
	}

	for row := 2; row <= maxGroups; row++ {
		for col := row * m; col <= n; col++ {
			for i := m; i <= col-m; i++ {
				ret[row][col] += ret[row-1][col-i]
			}
		}
	}

	for col := 1; col <= n; col++ {
		ret[maxGroups+1][col] = ret[1][col]
		for j := 2; j <= maxGroups; j++ {
			ret[maxGroups+1][col] += ret[j][col]
		}
	}

	return ret[maxGroups+1][n]
}
