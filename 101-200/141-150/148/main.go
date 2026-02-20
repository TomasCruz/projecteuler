package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 148; Exploring Pascal's Triangle
We can easily verify that none of the entries in the first seven rows of Pascal's triangle are divisible by 7:

 1
 1  1
 1  2  1
 1  3  3  1
 1  4  6  4  1
 1  5 10 10  5  1
 1  6 15 20 15  6  1

However, if we check the first one hundred rows, we will find that only 2361 of the 5050 entries are NOT divisible by 7.

Find the number of entries which are NOT divisible by 7 in the first one billion (10^9) rows of Pascal's triangle.
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
		limit = 9
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	power := 1
	for i := 0; i < limit; i++ {
		power *= 10
	}

	sum := countNonDivisibles(power)

	result = strconv.FormatInt(sum, 10)
	return
}

func base7(x int) []int64 {
	r := []int64{}
	for x > 0 {
		r = append(r, int64(x%7))
		x /= 7
	}

	for j, k := 0, len(r)-1; j < k; j, k = j+1, k-1 {
		r[j], r[k] = r[k], r[j]
	}

	return r
}

func countNonDivisibles(x int) int64 {
	digits := base7(x)
	digits = append([]int64{0}, digits...)

	powers := []int64{}
	p := int64(1)
	l := len(digits)
	for i := 1; i < l; i++ {
		powers = append(powers, p)
		p *= int64(28)
	}

	ret := int64(0)
	prev := int64(1)

	for i := 1; i < l; i++ {
		curr := prev * (digits[i-1] + 1)
		ret += curr * (digits[i] * (digits[i] + 1) / 2) * powers[l-i-1]
		prev = curr
	}

	return ret
}

/*
10 -> 40            13      28 + 2*6 = 40
100 -> 2361         202     3*28^2 + 3*3 = 2361
1000 -> 118335      2626    3*28^3 + 3*21*28^2 + 3*21*28 + 63*21 = 118335
                            (2)*28^3 + (2->3)*(6)*28^2 + [3*7]*(2)*28 + [3*7*3]*(6)
*/
