package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 71; Ordered Fractions
Consider the fraction, n/d, where n and d are positive integers. If n<d and HCF(n,d)=1, it is called a reduced proper fraction.
If we list the set of reduced proper fractions for d<=8 in ascending order of size, we get:
1/8,1/7,1/6,1/5,1/4,2/7,1/3,3/8,2/5,3/7,1/2,4/7,3/5,5/8,2/3,5/7,3/4,4/5,5/6,6/7,7/8
It can be seen that 2/5 is the fraction immediately to the left of 3/7.
By listing the set of reduced proper fractions for d<=1000000 in ascending order of size, find the numerator of the fraction
immediately to the left of 3/7.
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

	fs := make([]projecteuler.Fraction, limit+1)
	for d, m := 3, 3; d <= limit; d++ {
		x := (3 * d) / 7 // x/d < 3/7 => x < 3*d/7
		if m == 7 {
			m = 0
			x -= 1
		}

		fs[d] = projecteuler.NewFraction(int64(x), int64(d))
		m++
	}

	max := projecteuler.NewFraction(int64(1), int64(3))
	for d := 3; d <= limit; d++ {
		if max.Less(fs[d]) {
			max = fs[d]
		}
	}

	result = strconv.FormatInt(max.Num, 10)
	return
}
