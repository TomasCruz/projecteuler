package main

import (
	"sort"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 49; Prime permutations

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways:
(i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is
one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	primes, primeSet := projecteuler.PrimeSet(10000)
	primesCount := len(primes)

	for i := 0; i < primesCount; i++ {
		if primes[i] < 1000 {
			delete(primeSet, primes[i])
		}
	}

	seqFound := false
	for i := 0; i < primesCount && !seqFound; i++ {
		if _, ok := primeSet[primes[i]]; !ok {
			continue
		}

		digitOccurencies := digitOccurencies(primes[i])
		permutations := projecteuler.Permutations(4, func(args ...interface{}) bool { return false })
		permutatedPrimesFound := make(map[int]struct{})

		for _, perm := range permutations {
			num := 1000*int(digitOccurencies[perm[0]]) + 100*int(digitOccurencies[perm[1]]) +
				10*int(digitOccurencies[perm[2]]) + int(digitOccurencies[perm[3]])

			if _, ok := primeSet[num]; ok {
				delete(primeSet, num)
				permutatedPrimesFound[num] = struct{}{}
			}
		}

		if len(permutatedPrimesFound) > 2 {
			sequences := aritmeticSequence(permutatedPrimesFound)
			if sequences == nil {
				continue
			}

			for _, seq := range sequences {
				if seq[0] == 1487 && seq[1] == 4817 {
					continue
				}

				seqFound = true
				resultNum := 10000*seq[0] + seq[1]
				resultNum = 10000*resultNum + seq[2]
				result = strconv.FormatInt(int64(resultNum), 10)
				break
			}
		}
	}

	return
}

func digitOccurencies(x int) []byte {
	dn := projecteuler.NewDigitalNumber(x)
	return dn.Digits()
}

func aritmeticSequence(numbers map[int]struct{}) (sequences [][]int) {
	numCount := len(numbers)
	numSlice := make([]int, numCount)
	i := 0

	for n := range numbers {
		numSlice[i] = n
		i++
	}

	sort.Ints(numSlice)

	for i := 0; i < numCount-2; i++ {
		for j := i + 1; j < numCount-1; j++ {
			target := numSlice[j] - numSlice[i]

			for k := j + 1; k < numCount; k++ {
				if numSlice[k]-numSlice[j] == target {
					sequences = append(sequences, []int{numSlice[i], numSlice[j], numSlice[k]})
				}
			}
		}
	}

	return
}
