package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 45; Triangular, pentagonal, and hexagonal

Triangle, pentagonal, and hexagonal numbers are generated by the following formulae:
Triangle 	  	Tn=n(n+1)/2			1, 3, 6, 10, 15, ...
Pentagonal	 	Pn=n(3n−1)/2		1, 5, 12, 22, 35, ...
Hexagonal 	  	Hn=n(2n−1)			1, 6, 15, 28, 45, ...

It can be verified that T285 = P165 = H143 = 40755.
Find the next triangle number that is also pentagonal and hexagonal.
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
		limit = 1000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	roots := projecteuler.ReverseSquares(limit)

	numberFound := 0
	for i := 1; i < limit; i++ {
		hex := i * (2*i - 1)
		if projecteuler.IsPentagonal(hex, roots) && projecteuler.IsTriangular(hex, roots) {
			numberFound++
			if numberFound == 3 {
				result = strconv.FormatInt(int64(hex), 10)
				break
			}
		}
	}

	return
}
