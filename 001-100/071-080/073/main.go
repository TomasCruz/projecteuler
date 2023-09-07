package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 73; Counting Fractions in a Range
Consider the fraction, n/d, where n and d are positive integers. If n<d and HCF(n,d)=1, it is called a reduced proper fraction.
If we list the set of reduced proper fractions for d<=8 in ascending order of size, we get:
1/8,1/7,1/6,1/5,1/4,2/7,1/3,3/8,2/5,3/7,1/2,4/7,3/5,5/8,2/3,5/7,3/4,4/5,5/6,6/7,7/8
It can be seen that there are 3 fractions between 1/3 and 1/2.
How many fractions lie between 1/3 and 1/2 in the sorted set of reduced proper fractions for d <= 12000?
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
		limit = 12000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	count := int64(0)
	for d, m2 := 5, 1; d <= limit; d++ {
		// 1/3 < n/d < 1/2 => d/3 < n < d/2
		dHalf := d / 2
		if m2 == 2 {
			m2 = 0
			dHalf--
		}

		for i := d/3 + 1; i <= dHalf; i++ {
			if projecteuler.GCD(int64(i), int64(d)) == 1 {
				count++
			}
		}

		m2++
	}

	result = strconv.FormatInt(count, 10)
	return
}
