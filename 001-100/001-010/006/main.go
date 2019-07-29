package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 6; Sum square difference

The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers
and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
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
	resultInt := littleGausSum(limit)*littleGausSum(limit) - squarePyramidalNumber(limit)
	result = strconv.Itoa(resultInt)

	return
}

func squarePyramidalNumber(limit int) int {
	return limit * (limit + 1) * (2*limit + 1) / 6
}

func littleGausSum(limit int) int {
	return limit * (limit + 1) / 2
}
