package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 112; Bouncy Numbers
Working from left-to-right if no digit is exceeded by the digit to its left it is called an increasing number; for example, 134468.
Similarly if no digit is exceeded by the digit to its right it is called a decreasing number; for example, 66420.

We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number; for example, 155349.

Clearly there cannot be any bouncy numbers below one-hundred, but just over half of the numbers below one-thousand (525) are bouncy.
In fact, the least number for which the proportion of bouncy numbers first reaches 50% is 538.

Surprisingly, bouncy numbers become more and more common and by the time we reach 21780 the proportion of bouncy numbers is equal to 90%.

Find the least number for which the proportion of bouncy numbers is exactly 99%.
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
		limit = 99
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	incMatrix := calcIncMatrix(limit)
	decMatrix := calcDecMatrix(limit)

	nonBouncy := make([]uint64, limit+1)
	for i := 3; i <= limit; i++ {
		nonBouncy[i] = incMatrix[i][1] + decMatrix[i][9] - 10
	}

	for i := 4; i <= limit; i++ {
		nonBouncy[i] += nonBouncy[i-1]
	}

	for i := 3; i <= limit; i++ {
		nonBouncy[i] += uint64(99)
	}

	bouncy := make([]uint64, limit+1)
	power := uint64(100)
	i := 3
	for ; i <= limit; i++ {
		power *= 10
		bouncy[i] = power - nonBouncy[i] - 1
		percent := 100.0 * float64(bouncy[i]) / float64(power-1)
		if percent > float64(limit) {
			break
		}
	}

	i--
	power /= 10
	bnc := bouncy[i]
	j := power
	for ; ; j++ {
		if bounciness(limit, j, &bnc) {
			break
		}
	}

	result = strconv.FormatUint(j, 10)
	return
}

func bounciness(limit int, total uint64, bouncy *uint64) bool {
	digits := getDigs(total)

	if !inc(digits) && !dec(digits) {
		*bouncy++
		return cmp(*bouncy, total, limit, 100) == 0
	}

	return false
}

func getDigs(total uint64) []int {
	ret := []int{}

	for total != 0 {
		ret = append(ret, int(total%10))
		total /= 10
	}

	for l, r := 0, len(ret)-1; l < r; l, r = l+1, r-1 {
		ret[l], ret[r] = ret[r], ret[l]
	}

	return ret
}

func inc(digits []int) bool {
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}
	}

	return true
}

func dec(digits []int) bool {
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] < digits[i+1] {
			return false
		}
	}

	return true
}

func cmp(a, b uint64, c, d int) int {
	// a/b < c/d => a*d < b*c
	l, r := a*uint64(d), b*uint64(c)
	if l < r {
		return -1
	} else if l > r {
		return 1
	}

	return 0
}

func calcIncMatrix(limit int) [][]uint64 {
	ret := make([][]uint64, limit+1)

	for i := 1; i <= limit; i++ {
		ret[i] = make([]uint64, 10)
	}

	for i := 1; i < 10; i++ {
		ret[1][i] = uint64(10 - i)
	}

	for i := 2; i <= limit; i++ {
		ret[i][9] = 1
	}

	for i := 2; i <= limit; i++ {
		for j := 9; j > 1; j-- {
			ret[i][j-1] = ret[i][j] + ret[i-1][j-1]
		}
	}

	return ret
}

func calcDecMatrix(limit int) [][]uint64 {
	ret := make([][]uint64, limit+1)

	for i := 1; i <= limit; i++ {
		ret[i] = make([]uint64, 10)
	}

	for i := 0; i < 10; i++ {
		ret[1][i] = uint64(1 + i)
	}

	for i := 2; i <= limit; i++ {
		ret[i][0] = 1
	}

	for i := 2; i <= limit; i++ {
		for j := 0; j < 9; j++ {
			ret[i][j+1] = ret[i][j] + ret[i-1][j+1]
		}
	}

	return ret
}
