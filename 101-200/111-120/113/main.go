package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 113; Non-bouncy Numbers
Working from left-to-right if no digit is exceeded by the digit to its left it is called an increasing number; for example, 134468.

Similarly if no digit is exceeded by the digit to its right it is called a decreasing number; for example, 66420.

We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number; for example, 155349.

As n increases, the proportion of bouncy numbers below n increases such that there are only 12951 numbers below one-million that are not bouncy and only 277032 non-bouncy numbers below 10^10.

How many numbers below a googol (10^100) are not bouncy?
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

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	incMatrix := calcIncMatrix(limit)
	decMatrix := calcDecMatrix(limit)

	nonBouncy := make([]uint64, limit+1)
	for i := 3; i <= limit; i++ {
		nonBouncy[i] = incMatrix[i][1] + decMatrix[i][9] - 10
	}

	for i := 4; i <= limit; i++ {
		nonBouncy[i] += nonBouncy[i-1]
	}

	for i := 3; i <= limit; i++ {
		nonBouncy[i] += uint64(99)
	}

	result = strconv.FormatUint(nonBouncy[limit], 10)
	return
}

func calcIncMatrix(limit int) [][]uint64 {
	ret := make([][]uint64, limit+1)

	for i := 1; i <= limit; i++ {
		ret[i] = make([]uint64, 10)
	}

	for i := 1; i < 10; i++ {
		ret[1][i] = uint64(10 - i)
	}

	for i := 2; i <= limit; i++ {
		ret[i][9] = 1
	}

	for i := 2; i <= limit; i++ {
		for j := 9; j > 1; j-- {
			ret[i][j-1] = ret[i][j] + ret[i-1][j-1]
		}
	}

	return ret
}

func calcDecMatrix(limit int) [][]uint64 {
	ret := make([][]uint64, limit+1)

	for i := 1; i <= limit; i++ {
		ret[i] = make([]uint64, 10)
	}

	for i := 0; i < 10; i++ {
		ret[1][i] = uint64(1 + i)
	}

	for i := 2; i <= limit; i++ {
		ret[i][0] = 1
	}

	for i := 2; i <= limit; i++ {
		for j := 0; j < 9; j++ {
			ret[i][j+1] = ret[i][j] + ret[i-1][j+1]
		}
	}

	return ret
}
