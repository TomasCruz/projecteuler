package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 111; Primes with Runs
Considering 4-digit primes containing repeated digits it is clear that they cannot all be the same: 1111 is divisible by 11, 2222 is divisible by 22, and so on.
But there are nine 4-digit primes containing three ones:
1117, 1151, 1171, 1181, 1511, 1811, 2111, 4111, 8111.

We shall say that M(n, d) represents the maximum number of repeated digits for an n-digit prime where d is the repeated digit,
N(n, d) represents the number of such primes, and S(n, d) represents the sum of these primes.

So M(4, 1) = 3 is the maximum number of repeated digits for a 4-digit prime where one is the repeated digit, there are N(4, 1) = 9 such primes,
and the sum of these primes is S(4, 1) = 22275. It turns out that for d = 0, it is only possible to have M(4, 0) = 2 repeated digits,
but there are N(4, 0) = 13 such cases.

In the same way we obtain the following results for 4-digit primes.
Digit		M(4, d)		N(4, d)		S(4, d)
0			2				13		67061
1			3				9		22275
2			3				1		2221
3			3				12		46214
4			3				2		8888
5			3				1		5557
6			3				1		6661
7			3				9		57863
8			3				1		8887
9			3				7		48073

For d = 0 to 9, the sum of all S(4, d) is 273700.
Find the sum of all S(10, d).
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
		limit = 10
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	orderedMasks := generateMasks(limit)
	primes := projecteuler.Primes(int(math.Pow10(limit/2)), nil)

	digitMaps := make([]map[int64]struct{}, 10)
	for digit, digitCountSlice := range orderedMasks {
		for count := len(digitCountSlice) - 1; count > 1; count-- {
			x := genPrimesFromMasks(limit, digit, digitCountSlice[count], primes)
			if len(x) > 0 {
				digitMaps[digit] = x
				break
			}
		}
	}

	primeSums := make([]int64, 10)
	for digit, m := range digitMaps {
		sum := int64(0)
		for x := range m {
			sum += x
		}
		primeSums[digit] = sum
	}

	totalSum := int64(0)
	for _, s := range primeSums {
		totalSum += s
	}

	result = strconv.FormatInt(totalSum, 10)
	return
}

func genPrimesFromMasks(limit, digit int, masks map[string]struct{}, primes []int) map[int64]struct{} {
	ret := map[int64]struct{}{}

	for k := range masks {
		cm := genPrimesFromCurrentMask(limit, digit, k, primes)
		addToMap(ret, cm)
	}

	return ret
}

func genPrimesFromCurrentMask(limit, digit int, mask string, primes []int) map[int64]struct{} {
	ret := map[int64]struct{}{}
	currNFCM := newNFCM(limit, digit, mask, primes)
	currNFCM.genPrimesRec(0, int64(0), ret)
	return ret
}

func addToMap(a, b map[int64]struct{}) {
	for k := range b {
		a[k] = struct{}{}
	}
}
