package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 125; Palindromic Sums
The palindromic number 595 is interesting because it can be written as the sum of consecutive squares: 6^2 + 7^2 + 8^2 + 9^2 + 10^2 + 11^2 + 12^2.

There are exactly eleven palindromes below one-thousand that can be written as consecutive square sums, and the sum of these palindromes is 4164.
Note that 1 = 0^2 + 1^2 has not been included as this problem is concerned with the squares of positive integers.

Find the sum of all the numbers less than 10^8 that are both palindromic and can be written as the sum of consecutive squares.
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
		limit = 8
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	power := 1
	for range limit {
		power *= 10
	}

	squareSums := calcSquareSums(power)
	palindroms := calcPalindroms(limit)

	pals := map[int]struct{}{}
	for i := 2; i < len(squareSums); i++ {
		for j := 1; j < len(squareSums[i]); j++ {
			if _, present := palindroms[squareSums[i][j]]; present {
				pals[squareSums[i][j]] = struct{}{}
			}
		}
	}

	sum := int64(0)
	for x := range pals {
		sum += int64(x)
	}

	result = strconv.FormatInt(sum, 10)
	return
}

func calcSquareSums(power int) [][]int {
	squareSums := [][]int{{}, {0}}
	for i := 1; ; i++ {
		sq := i * i
		if sq > power {
			break
		}

		squareSums[1] = append(squareSums[1], sq)
	}

	for i := 2; i <= len(squareSums[1]); i++ {
		newLen := len(squareSums[i-1]) - 1
		if newLen == 0 {
			break
		}

		squareSums = append(squareSums, make([]int, 1, newLen))
		for j := 1; j < newLen; j++ {
			x := squareSums[1][j] + squareSums[i-1][j+1]
			if x > power {
				break
			}

			squareSums[i] = append(squareSums[i], x)
		}
	}

	return squareSums
}

func calcPalindroms(limit int) map[int]struct{} {
	ret := map[int]struct{}{}

	power := 1
	nextPower := 1

	half := limit / 2
	if limit%2 == 1 {
		half++
	}

	for i := 1; i <= half; i++ {
		nextPower = power * 10
		for k := power; k < nextPower; k++ {
			ret[mirror(k, power, true)] = struct{}{}
			ret[mirror(k, nextPower, false)] = struct{}{}
		}
		power *= 10
	}

	return ret
}

func mirror(k, power int, odd bool) int {
	digits := []int{}

	n := k
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	n = k * power
	if odd {
		digits = digits[1:]
	}

	for _, d := range digits {
		power /= 10
		n += power * d
	}

	return n
}
