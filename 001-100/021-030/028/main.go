package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 28; Number spiral diagonals

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
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
		limit = 1001
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	sum := 1 // central
	for i := 3; i <= limit; i += 2 {
		sum += i*(i-1) + 1
		sum += i * i
		sum += (i-2)*(i-1) + 1
		sum += (i-1)*(i-1) + 1
	}

	result = strconv.Itoa(sum)
	return
}
