package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 78; Coin Partitions
Let p(n) represent the number of different ways in which coins can be separated into piles.
For example, five coins can be separated into piles in exactly seven different ways, so p(5)=7.
OOOOO
OOOO   O
OOO   OO
OOO   O   O
OO   OO   O
OO   O   O   O
O   O   O   O   O

Find the least value of n for which p(n) is divisible by one million.
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
	p := partitions(limit)
	result = strconv.Itoa(len(p) - 1)

	return
}

func partitions(limit int) (partitions []int) {
	partitions = make([]int, 0, limit+1)
	partitions = append(partitions, 1)
	partitions = append(partitions, 1)
	partitions = append(partitions, 2)

	for i := 3; i < limit; i++ {
		sum := 0
		sign := 1

		for k := 1; k <= i; k++ {
			aIndex := i - k*(3*k-1)/2
			a := 0
			if aIndex >= 0 {
				a = partitions[aIndex]
			}

			bIndex := i - k*(3*k+1)/2
			b := 0
			if bIndex >= 0 {
				b = partitions[bIndex]
			}

			term := sign * (a + b) % limit
			sum = (sum + term) % limit
			if b == 0 {
				break
			}

			sign = -1 * sign
		}

		partitions = append(partitions, sum)
		if sum == 0 {
			break
		}
	}

	return
}
