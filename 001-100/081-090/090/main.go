package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 90; Cube Digit Pairs
Each of the six faces on a cube has a different digit (0 to 9) written on it; the same is done to a second cube.
By placing the two cubes side-by-side in different positions we can form a variety of 2-digit numbers.

For example, the square number 64 could be formed:

In fact, by carefully choosing the digits on both cubes it is possible to display all of the square numbers below one-hundred:
01, 04, 09, 16, 25, 36, 49, 64, and 81.

For example, one way this can be achieved is by placing {0, 5, 6, 7, 8, 9} on one cube and {1, 2, 3, 4, 8, 9} on the other cube.

However, for this problem we shall allow the 6 or 9 to be turned upside-down so that an arrangement like {0, 5, 6, 7, 8, 9} and {1, 2, 3, 4, 6, 7}
allows for all nine square numbers to be displayed; otherwise it would be impossible to obtain 09.

In determining a distinct arrangement we are interested in the digits on each cube, not the order.
{1, 2, 3, 4, 5, 6} is equivalent to {3, 6, 4, 1, 2, 5}
{1, 2, 3, 4, 5, 6} is distinct from {1, 2, 3, 4, 5, 9}
But because we are allowing 6 and 9 to be reversed, the two distinct sets in the last example both represent the extended set
{1, 2, 3, 4, 5, 6, 9} for the purpose of forming 2-digit numbers.

How many distinct arrangements of the two cubes allow for all of the square numbers to be displayed?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	allCombos, _ := projecteuler.Combinations(10, 6, nil)

	validArrs := []arrangement{}
	for _, currCombo := range allCombos {
		arr := makeArr(currCombo)

		if arr.isValid() {
			validArrs = append(validArrs, arr)
		}
	}

	pairArrs := map[arrangement]map[arrangement]struct{}{}
	for i := 0; i < len(validArrs); i++ {
		for j := 0; j < len(validArrs); j++ {
			if validArrs[i].isComplementValid(validArrs[j]) {
				addArrangementPairSet(validArrs[i], validArrs[j], pairArrs)
			}
		}
	}

	res := 0
	for arr := range pairArrs {
		v := 1
		if arr.hasDigit[6] == 1 && arr.value>>12 == 7 {
			v *= 2
		}
		for k := range pairArrs[arr] {
			currV := v
			if k.hasDigit[6] == 1 && k.value>>12 == 7 {
				currV *= 2
			}
			res += currV
		}
	}

	result = strconv.Itoa(res)
	return
}

func addArrangementPairSet(l, r arrangement, pairArrs map[arrangement]map[arrangement]struct{}) {
	if l.value < r.value {
		if _, exists := pairArrs[l]; !exists {
			pairArrs[l] = map[arrangement]struct{}{}
		}
		pairArrs[l][r] = struct{}{}
	} else {
		if _, exists := pairArrs[r]; !exists {
			pairArrs[r] = map[arrangement]struct{}{}
		}
		pairArrs[r][l] = struct{}{}
	}
}
