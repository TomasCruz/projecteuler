package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 92; Square Digit Chains
A number chain is created by continuously adding the square of the digits in a number
to form a new number until it has been seen before.
For example,

44 -> 32 -> 13 -> 10 -> 1 -> 1
85 -> 89 -> 145 -> 42 -> 20 -> 4 -> 16 -> 37 -> 58 -> 89

Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop. What is most amazing is
that EVERY starting number will eventually arrive at 1 or 89.
How many starting numbers below ten million will arrive at 89?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	res := 0
	for i := 2; i < 10000000; i++ {
		d := i

		for {
			next := 0
			for d > 0 {
				curr := d % 10
				next += curr * curr
				d /= 10
			}

			if next == 89 {
				res++
				break
			} else if next == 1 {
				break
			}

			d = next
		}
	}

	result = strconv.Itoa(res)
	return
}
