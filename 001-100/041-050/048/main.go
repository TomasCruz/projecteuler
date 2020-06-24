package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 48; Self powers

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.
Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
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
		limit = 1000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	magicNumber := int64(10000000000)

	var sum int64
	for i := 1; i <= limit; i++ {
		sum = projecteuler.ModuloSum(sum, projecteuler.ModuloSelfPower(i, magicNumber), magicNumber)
	}

	result = strconv.FormatInt(sum, 10)
	return
}
