package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 40; Champernowne's constant

An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
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
		limit = 6
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	// digitNumbers are counts of numbers with n digits, i.e. of numbers [10^(n-1), 10^n),
	// i.e. digitNumbers[n] = 10^n - 10^(n-1)
	// digitNumbers[0]=0, digitNumbers[1]=9 (10^0==1..10), digitNumbers[2]=90 (10^1..10^2), digitNumbers[3]=900
	digitNumbers := []int{0, 9, 90, 900, 9000, 90000, 900000}

	// digitsUsed[n] is count of digits used for numbers of up to n digits,
	// i.e. digitsUsed[2] = 0*0 + 9*1 + 90*2 = 189
	digitsUsed := make([]int, len(digitNumbers))
	for i := 1; i < len(digitNumbers); i++ {
		digitsUsed[i] = digitsUsed[i-1] + digitNumbers[i]*i
	}

	pow := 2
	powered := 100
	n := 1
	for pow <= limit {
		n *= d(powered, digitNumbers, digitsUsed)
		pow++
		powered *= 10
	}

	result = strconv.Itoa(n)
	return
}

func d(n int, digitNumbers []int, digitsUsed []int) (digit int) {
	i := 1

	for n > digitsUsed[i] {
		i++
	}

	// i is now number of digits of number containing d(n)
	n -= digitsUsed[i-1]
	number := digitNumbers[i]/9 + n/i
	n -= (n / i) * i
	dn := projecteuler.NewDigitalNumber(number)

	digit = int(dn.Digits()[i-n])
	return
}
