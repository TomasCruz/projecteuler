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

	m := projecteuler.GeneratePartitionMatrix(limit)
	sum := 0
	for i := 1; i < limit; i++ {
		sum += m[limit][i]
	}

	result = strconv.Itoa(sum)
	return
}
