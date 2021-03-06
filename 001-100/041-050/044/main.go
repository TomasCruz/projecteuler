package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 44; Pentagon numbers

Pentagonal numbers are generated by the formula, Pn=n(3n−1)/2. The first ten pentagonal numbers are:

1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...

It can be seen that P4 + P7 = 22 + 70 = 92 = P8. However, their difference, 70 − 22 = 48, is not pentagonal.
Find the pair of pentagonal numbers, Pj and Pk, for which their sum and difference are pentagonal and
D = |Pk − Pj| is minimised; what is the value of D?
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
		limit = 20000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	roots := projecteuler.ReverseSquares(limit)

	minDist := limit * limit
	notFound := true
	for i := 1; i < limit && notFound; i++ {
		for j := 1; i+j < limit; j++ {
			dist := j * (3*j + 6*i - 1) / 2
			if projecteuler.IsPentagonal(dist, roots) {
				sum := dist + i*(6*i-2)/2
				if projecteuler.IsPentagonal(sum, roots) {
					if dist < minDist {
						minDist = dist
						notFound = false
					}
					break
				}
			}
		}
	}

	result = strconv.FormatInt(int64(minDist), 10)
	return
}
