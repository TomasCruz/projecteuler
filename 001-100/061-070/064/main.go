package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 64; Odd period square roots

see ./problem064.pdf
*/

func main() {
	projecteuler.Timed(calc, 10000)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(limit+1, nil)
	sqMap := projecteuler.ReverseSquares(limit)
	r := 0

	for x := 2; x <= limit; x++ {
		if _, ok := sqMap[x]; ok {
			continue
		}

		c := projecteuler.MakeContinuedFraction(x, primes)
		if c.Period()%2 == 1 {
			r++
			//fmt.Println(c.String())
		}
	}

	result = strconv.FormatInt(int64(r), 10)
	return
}
