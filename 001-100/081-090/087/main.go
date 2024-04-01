package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 87; Prime Power Triples
The smallest number expressible as the sum of a prime square, prime cube, and prime fourth power is 28. In fact,
there are exactly four numbers below fifty that can be expressed in such a way:
28 = 2^2 + 2^3 + 2^4
33 = 3^2 + 2^3 + 2^4
49 = 5^2 + 2^3 + 2^4
47 = 2^2 + 3^3 + 2^4

How many numbers below fifty million can be expressed as the sum of a prime square, prime cube, and prime fourth power?
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
		limit = 50000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	sqLimit := int(math.Ceil(math.Sqrt(float64(limit))))
	primes := projecteuler.Primes(sqLimit, nil)

	var squares []int
	for i := 0; i < len(primes); i++ {
		v := primes[i] * primes[i]
		if v > limit {
			break
		}
		squares = append(squares, v)
	}

	var cubes []int
	for i := 0; i < len(primes); i++ {
		v := primes[i] * primes[i] * primes[i]
		if v > limit {
			break
		}
		cubes = append(cubes, v)
	}

	var fourths []int
	for i := 0; i < len(primes); i++ {
		v := primes[i] * primes[i] * primes[i] * primes[i]
		if v > limit {
			break
		}
		fourths = append(fourths, v)
	}

	resSet := map[int]struct{}{}
	for i := 0; i < len(fourths); i++ {
		for j := 0; j < len(cubes); j++ {
			sum := fourths[i] + cubes[j]
			if sum > limit {
				break
			}

			for k := 0; k < len(squares); k++ {
				num := sum + squares[k]
				if num > limit {
					break
				}

				resSet[num] = struct{}{}
			}
		}
	}

	result = strconv.Itoa(len(resSet))
	return
}
