package main

import (
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 20; Factorial digit sum

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
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

	resultBig := big.NewInt(int64(2))
	for i := 3; i <= limit; i++ {
		resultBig.Mul(resultBig, big.NewInt(int64(i)))
	}

	digitSum := 0
	str := resultBig.Text(10)

	for i := 0; i < len(str); i++ {
		digitSum += int(str[i] - '0')
	}

	result = strconv.Itoa(digitSum)
	return
}
