package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 77; Prime Summations
It is possible to write ten as the sum of primes in exactly five different ways:
7+3
5+5
5+3+2
3+3+2+2
2+2+2+2+2
What is the first value which can be written as the sum of primes in over five thousand different ways?
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
		limit = 5000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(limit, nil)
	m := generatePartitionMatrix(limit, primes)

	i := 0
	for ; i <= limit; i++ {
		sum := 0
		for j := 0; j < len(m[0]); j++ {
			sum += m[i][j]
		}
		if sum > limit {
			break
		}
	}

	result = strconv.Itoa(i)
	return
}

func generatePartitionMatrix(limit int, primes []int) [][]int {
	i := 0
	for ; i < len(primes) && primes[i] <= limit; i++ {
	}
	colCount := i

	m := make([][]int, limit+1)
	for i := 0; i <= limit; i++ {
		m[i] = make([]int, colCount)
	}

	m[0][0] = 1
	for i := 2; i <= limit; i++ {
		if i%2 == 0 {
			m[i][0] = 1
		} else {
			m[i][0] = 0
		}
	}

	for i := 3; i <= limit; i++ {
		for j := 1; j < colCount; j++ {
			sum := 0
			for k := 1; k*primes[j] <= i; k++ {
				for l := 0; l < j; l++ {
					sum += m[i-k*primes[j]][l]
				}
			}
			m[i][j] = sum
		}
	}

	return m
}
