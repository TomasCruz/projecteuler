package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 76; Counting Summations
It is possible to write five as a sum in exactly six different ways:
4+1
3+2
3+1+1
2+2+1
2+1+1+1
1+1+1+1+1
How many different ways can one hundred be written as a sum of at least two positive integers?
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

	m := generatePartitionMatrix(limit)
	sum := 0
	for i := 1; i < limit; i++ {
		sum += m[limit][i]
	}

	result = strconv.Itoa(sum)
	return
}

func generatePartitionMatrix(limit int) [][]int {
	// init
	m := make([][]int, limit+1)
	for i := 0; i <= limit; i++ {
		m[i] = make([]int, limit+1)
	}

	m[1][1] = 1
	for i := 2; i <= limit; i++ {
		m[i][1] = 1
		m[i][2] = i / 2
		m[i][i-1] = 1
		m[i][i] = 1
	}

	for i := 5; i <= limit; i++ {
		for j := 3; j < i-1; j++ {
			sum := 0
			for k := 1; k <= j; k++ {
				sum += m[i-j][k]
			}
			m[i][j] = sum
		}
	}

	return m
}
