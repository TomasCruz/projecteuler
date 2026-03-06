package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 117; Red, Green, and Blue Tiles
Using a combination of grey square tiles and oblong tiles chosen from:
red tiles (measuring two units), green tiles (measuring three units), and blue tiles (measuring four units),
it is possible to tile a row measuring five units in length in exactly fifteen different ways.

How many ways can a row measuring fifty units in length be tiled?

NOTE: This is related to Problem 116.
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

	total := make([]uint64, 4, limit+1)
	total[0] = uint64(1)
	total[1] = uint64(1)
	total[2] = uint64(2)
	total[3] = uint64(4)

	total = ways(limit, total)

	result = strconv.FormatUint(total[limit], 10)
	return
}

func ways(n int, total []uint64) []uint64 {
	if n < len(total) {
		return total
	}

	total = ways(n-1, total)
	if n > 3 {
		total = ways(n-4, total)
	}
	if n > 2 {
		total = ways(n-3, total)
	}
	if n > 1 {
		total = ways(n-2, total)
	}

	if n == len(total) {
		sum := total[n-1] + total[n-2] + total[n-3] + total[n-4]
		total = append(total, sum)
	}

	return total
}
