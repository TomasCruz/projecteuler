package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 35; Circular primes

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
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

type digitalPrime struct {
	used bool
	projecteuler.DigitalNumber
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(limit, nil)
	digitalPrimeSet := make(map[int]*digitalPrime)

	for i := 0; i < len(primes); i++ {
		digitalPrimeSet[primes[i]] = &digitalPrime{DigitalNumber: projecteuler.NewDigitalNumber(primes[i])}
	}

	circularPrimesCount := 0
	for _, v := range digitalPrimeSet {
		if v.used {
			continue
		}
		v.used = true

		circular := true
		count := 1
		digCount := v.DigitCount()
		byteMatrix := rotations(v.Digits())

		for i := 1; i < digCount; i++ {
			curr := projecteuler.NumberFromDigits(byteMatrix[i])
			if digPrime, ok := digitalPrimeSet[curr]; !ok {
				circular = false
				break
			} else if !digPrime.used {
				count++
				digPrime.used = true
			}
		}

		if circular {
			circularPrimesCount += count
		}
	}

	result = strconv.Itoa(circularPrimesCount)
	return
}

func rotations(digits []byte) (digitRotations [][]byte) {
	length := len(digits)
	digitRotations = make([][]byte, length)
	digitRotations[0] = make([]byte, length)
	copy(digitRotations[0], digits)

	for i := 1; i < length; i++ {
		digitRotations[i] = make([]byte, length)
		first := digitRotations[i-1][0]

		for j := 1; j < length; j++ {
			digitRotations[i][j-1] = digitRotations[i-1][j]
		}

		digitRotations[i][length-1] = first
	}

	return
}
