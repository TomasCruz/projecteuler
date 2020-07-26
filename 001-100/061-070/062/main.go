package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 62; Cubic permutations

The cube, 41063625 (345^3), can be permuted to produce two other cubes: 56623104 (384^3) and 66430125 (405^3).
In fact, 41063625 is the smallest cube which has exactly three permutations of its digits which are also cube.

Find the smallest cube for which exactly five permutations of its digits are cube.
*/

func main() {
	var cubicPermutationsCount int

	if len(os.Args) > 1 {
		cubicPermutationsCount64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		cubicPermutationsCount = int(cubicPermutationsCount64)
	} else {
		cubicPermutationsCount = 5
	}

	projecteuler.Timed(calc, cubicPermutationsCount)
}

func calc(args ...interface{}) (result string, err error) {
	// 10^(n-1) <= n digit cube < 10^n
	// n-1 <= log(x) < n
	// 10^[(n-1)/3] <= x < 10^[n/3]

	cubicPermutationsCount := args[0].(int)
	weights := fillWeights()

	for digitCount := 3; ; digitCount++ {
		start := math.Floor(math.Pow(10, float64(digitCount-1)/3.0))
		end := math.Floor(math.Pow(10, float64(digitCount)/3.0))
		masks := make(map[int]map[int64]struct{})
		maskMin := make(map[int]int64)

		for i := start; i < end; i++ {
			cube := int64(i) * int64(i) * int64(i)
			currMask := getMask(cube, weights)

			if _, ok := masks[currMask]; !ok {
				masks[currMask] = make(map[int64]struct{})
				maskMin[currMask] = cube
			}
			masks[currMask][cube] = struct{}{}
		}

		res := int64(math.MaxInt64)
		for k, v := range masks {
			if len(v) == cubicPermutationsCount && maskMin[k] < res {
				res = maskMin[k]
			}
		}

		if res != int64(math.MaxInt64) {
			result = strconv.FormatInt(res, 10)
			return
		}
	}
}

func getMask(cube int64, weights []int) int {
	bi, _ := projecteuler.MakeBigInt(strconv.FormatInt(cube, 10))
	digits := []byte(bi.String())

	mask := 0
	for i := 0; i < len(digits); i++ {
		mask += weights[digits[i]-'0']
	}

	return mask
}

func fillWeights() []int {
	weights := make([]int, 10)
	weights[0] = 0x1
	for i := 1; i < 10; i++ {
		weights[i] = weights[i-1] * 8
	}

	return weights
}
