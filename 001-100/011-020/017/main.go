package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 17; Number letter counts

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters
used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains
23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in
compliance with British usage.
*/

func main() {
	projecteuler.TimedStr(calc)
}

func calc(args ...interface{}) (result string, err error) {
	numMap := buildMap()

	// 1-99
	sumBH := sumBelowHundred(numMap)
	sum := sumBH

	// 100 - 999
	for i := 1; i < 10; i++ {
		sum += numMap[i] + numMap[100] + (numMap[i]+numMap[100]+3)*99 + sumBH
	}

	// 1000
	sum += numMap[1] + numMap[1000]

	result = strconv.Itoa(sum)
	return
}

func buildMap() (numMap map[int]int) {
	numMap = make(map[int]int)
	numMap[1] = 3
	numMap[2] = 3
	numMap[3] = 5
	numMap[4] = 4
	numMap[5] = 4
	numMap[6] = 3
	numMap[7] = 5
	numMap[8] = 5
	numMap[9] = 4
	numMap[10] = 3
	numMap[11] = 6
	numMap[12] = 6
	numMap[13] = 8
	numMap[14] = 8
	numMap[15] = 7
	numMap[16] = 7
	numMap[17] = 9
	numMap[18] = 8
	numMap[19] = 8
	numMap[20] = 6
	numMap[30] = 6
	numMap[40] = 5
	numMap[50] = 5
	numMap[60] = 5
	numMap[70] = 7
	numMap[80] = 6
	numMap[90] = 6
	numMap[100] = 7
	numMap[1000] = 8

	return
}

func sumBelowHundred(numMap map[int]int) int {
	sumBT := sumBelowTen(numMap)
	sum := sumBT

	// 10-19
	for i := 10; i < 20; i++ {
		sum += numMap[i]
	}

	// 20-99
	for i := 2; i < 10; i++ {
		sum += numMap[i*10]*10 + sumBT
	}

	return sum
}

func sumBelowTen(numMap map[int]int) int {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += numMap[i]
	}

	return sum
}
