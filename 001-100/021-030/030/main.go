package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 30; Digit fifth powers

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

    1634 = 1^4 + 6^4 + 3^4 + 4^4
    8208 = 8^4 + 2^4 + 0^4 + 8^4
    9474 = 9^4 + 4^4 + 7^4 + 4^4

As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
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
		limit = 5
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	upperLimit := pow(10, limit+1)
	sum := 0

	for i := 10; i < upperLimit; i++ {
		if examineNumber(i, limit) {
			sum += i
		}
	}

	result = strconv.Itoa(sum)
	return
}

func examineNumber(x, limit int) bool {
	xCopy := x
	digitSum := 0

	for xCopy > 0 {
		digitSum += pow(xCopy%10, limit)
		xCopy /= 10
	}

	return x == digitSum
}

func pow(x, n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= x
	}

	return res
}
