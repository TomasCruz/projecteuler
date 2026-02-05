package main

import (
	"log"
	"math"
	"os"
	"slices"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 119; Digit Power Sum
The number 512 is interesting because it is equal to the sum of its digits raised to some power: 5 + 1 + 2 = 8, and 8^3 = 512.
Another example of a number with this property is 614656 = 28^4.

We shall define a_n to be the nth term of this sequence and insist that a number must contain at least two digits to have a sum.

You are given that a2 = 512 and a10 = 614656.

Find a30.
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
		limit = 30
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	a := []int64{}

	for digSum := 2; digSum < 100; digSum++ {
		logDS := math.Log10(float64(digSum))

		num := int64(1)
		for power := 1; ; power++ {
			numDigits := math.Round(logDS * float64(power))
			if numDigits > 18 {
				break
			}

			num *= int64(digSum)

			if int(digitSum(num)) == digSum && numDigits > 1 {
				a = append(a, num)
			}
		}
	}

	slices.Sort(a)

	result = strconv.FormatInt(a[limit-1], 10)
	return
}

func digitSum(i int64) int64 {
	sum := int64(0)

	for i > 0 {
		sum += i % 10
		i /= 10
	}

	return sum
}

/*
	81 = (8 + 1)^2, log(8+1)81 = 2
	512 = (5 +1 + 2)^3, log(8)512 = 3

	10^(numDigits-1) < digitSum^power < 10^numDigits
	numDigits-1 < power*log(digitSum) < numDigits
*/
