package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 64; Odd period square roots

see ./problem064.pdf
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
		limit = 10000
	}

	projecteuler.Timed(calc, limit)
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
