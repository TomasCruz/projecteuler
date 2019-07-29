package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 9; Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
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

	/*
		a < b < c
		a + b + c == 1000 (1)
		a^2 + b^2 == c^2  (2)
		-------------
		(a + b + c)^2 == 1000^2 == 10^6 // (1)^2
		a^2 + b^2 + c^2 + 2*(a*b + b*c + a*c) == 10^6 // replace from (2)
		2*(a^2 + b^2) + 2*(a*b + b*c + a*c) == 10^6 // /2
		(a^2 + b^2) + (a*b + b*c + a*c) == 5*10^5 // + a*b
		(a + b)^2 + c*(a + b) == 5*10^5 + a*b
		(a + b + c)*(a + b) == 5*10^5 + a*b // replace from (1)
		1000*(a + b) == 5*10^5 + a*b
		1000*b - a*b + 1000*a == 5*10^5
		b*(1000 - a) + 1000*a == 5*10^5
		(5*10^5 - 1000*a) / (1000-a) == b
		if (5*10^5 - 1000*a) % (1000-a) == 0, solution found
	*/

	resultInt := 0
	for a := 1; a < limit/3; a++ {
		if (limit*limit/2-limit*a)%(limit-a) == 0 {
			b := (limit*limit/2 - limit*a) / (limit - a)
			c := limit - a - b
			resultInt = a * b * c
			break
		}
	}

	result = strconv.Itoa(resultInt)
	return
}
