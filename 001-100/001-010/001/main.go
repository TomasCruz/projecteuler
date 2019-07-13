package main

import (
	"fmt"

	"github.com/TomasCruz/projecteuler"
)

/*
Multiples of 3 and 5
Problem 1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

func main() {
	//projecteuler.Timed(calc)
	projecteuler.Timed(calc, 1000)
}

func calc(args ...interface{}) {
	var limit int

	if len(args) == 0 {
		limit = 1000
	} else {
		limit = args[0].(int)
	}

	uLimit3 := biggestSmallerDivisibleBy(limit, 3)
	uLimit5 := biggestSmallerDivisibleBy(limit, 5)
	uLimit15 := biggestSmallerDivisibleBy(limit, 15)

	result := 3*littleGausSum(uLimit3/3) + 5*littleGausSum(uLimit5/5) - 15*littleGausSum(uLimit15/15)
	fmt.Println(result)
}

func littleGausSum(limit int) int {
	return limit * (limit + 1) / 2
}

func biggestSmallerDivisibleBy(limit int, divisor int) int {
	if divisor >= limit {
		return 0
	}

	var i int
	for i = limit - 1; i > 0; i-- {
		if i%divisor == 0 {
			break
		}
	}

	return i
}
