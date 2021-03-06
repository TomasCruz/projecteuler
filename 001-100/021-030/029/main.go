package main

import (
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Distinct powers
Problem 29

Consider all integer combinations of a^b for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

    2^2=4, 2^3=8, 2^4=16, 2^5=32
    3^2=9, 3^3=27, 3^4=81, 3^5=243
    4^2=16, 4^3=64, 4^4=256, 4^5=1024
    5^2=25, 5^3=125, 5^4=625, 5^5=3125

If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by a^b for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
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
		limit = 101
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	factorizationSlice := make([]map[int]int, 0)
	resultMap := make(map[int]map[int]struct{})
	primes := projecteuler.Primes(limit, nil)

	for a := 2; a < limit; a++ {
		factA, _ := projecteuler.Factorize(a, primes)
		currIndex, baseExponent := addFactorization(factA, &factorizationSlice)

		if _, ok := resultMap[currIndex]; !ok {
			resultMap[currIndex] = make(map[int]struct{})
		}

		for b := 2; b < limit; b++ {
			resultMap[currIndex][b*baseExponent] = struct{}{}
		}
	}

	sum := 0
	for _, v := range resultMap {
		sum += len(v)

	}

	result = strconv.Itoa(sum)
	return
}

type intSlice []int

func (a intSlice) Len() int           { return len(a) }
func (a intSlice) Less(i, j int) bool { return a[i] < a[j] }
func (a intSlice) Swap(i, j int)      { temp := a[i]; a[i], a[j] = a[j], temp }

func addFactorization(factors map[int]int, factorizationSlice *[]map[int]int) (index, baseExponent int) {
	powers := make([]int, len(factors))
	i := 0
	for _, v := range factors {
		powers[i] = v
		i++
	}
	sort.Sort(intSlice(powers))

	baseExponent = makePowersCoprime(powers)
	if baseExponent != 1 {
		for k, v := range factors {
			factors[k] = v / baseExponent
		}
	}

	for i := 0; i < len(*factorizationSlice); i++ {
		if projecteuler.CompareFactors(factors, (*factorizationSlice)[i]) {
			index = i
			return
		}
	}

	*factorizationSlice = append(*factorizationSlice, factors)
	index = len(*factorizationSlice) - 1

	return
}

func makePowersCoprime(powers []int) (baseExponent int) {
	length := len(powers)
	if length == 1 {
		baseExponent = powers[0]
		powers[0] = 1
		return
	}

	baseExponent = 1
	doBreak := false
	for i := 2; i <= powers[length-1]; i++ {
		var j int
		for j = 0; j < length && powers[j]%i == 0; j++ {
		}

		if j == length {
			baseExponent *= i
			for k := 0; k < length; k++ {
				powers[k] /= i
				if powers[k] == 1 {
					doBreak = true
				}
			}

			if doBreak {
				break
			}
		}
	}

	return
}
