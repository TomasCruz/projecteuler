package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 154; Exploring Pascal's Pyramid
A triangular pyramid is constructed using spherical balls so that each ball rests on exactly three balls of the next lower level.

Then, we calculate the number of paths leading from the apex to each position:
A path starts at the apex and progresses downwards to any of the three spheres directly below the current position.

Consequently, the number of paths to reach a certain position is the sum of the numbers immediately above it (depending on the position, there are up to three numbers above it).

The result is Pascal's pyramid and the numbers at each level n are the coefficients of the trinomial expansion (x + y + z)^n.
How many coefficients in the expansion of (x + y + z)^200000 are multiples of 10^12?
*/

func main() {
	var trinomialExponent int
	var multiplesCode string

	if len(os.Args) > 1 {
		trinomialExponent64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		trinomialExponent = int(trinomialExponent64)
	} else {
		trinomialExponent = 200000
	}

	if len(os.Args) > 2 {
		multiplesCode = os.Args[2]
	} else {
		multiplesCode = "2-12,5-12"
	}

	projecteuler.Timed(calc, trinomialExponent, multiplesCode)
}

func calc(args ...interface{}) (result string, err error) {
	trinomialExponent := args[0].(int)
	multiplesCode := args[1].(string)

	var multiples []pair
	multiples, err = parseMultiplesCode(multiplesCode)
	if err != nil {
		log.Fatal("bad multiplesCode arg")
	}

	// precalculate Legendre valuations
	l := len(multiples)
	v := make([][]int, l)
	for pIndex := range l {
		v[pIndex] = make([]int, trinomialExponent+1)
		for i := 1; i <= trinomialExponent; i++ {
			v[pIndex][i] = vp(multiples[pIndex].p, i)
		}
	}

	sum := int64(0)
	for i := 0; i <= trinomialExponent; i++ {
		restJK := trinomialExponent - i
		restJKHalf := (restJK + 1) / 2
		for j := restJKHalf; j <= min(restJK, i); j++ {
			k := restJK - j

			sum += examine(trinomialExponent, i, j, k, multiples, v)
		}
	}

	result = strconv.FormatInt(sum, 10)
	return
}

type pair struct {
	p, e int
}

func examine(trinomialExponent, i, j, k int, multiples []pair, v [][]int) int64 {
	l := len(multiples)
	for pIndex := range l {
		val := v[pIndex][trinomialExponent] - v[pIndex][i] - v[pIndex][j] - v[pIndex][k]
		if val < multiples[pIndex].e {
			return 0
		}
	}

	diff := 1
	if i != j {
		diff++
	}
	if j != k {
		diff++
	}

	mul := int64(1)
	switch diff {
	case 2:
		mul = 3
	case 3:
		mul = 6
	}

	return mul
}

// Legendre p-adic valuation (how many times p divides n!)
func vp(p, n int) int {
	sum := 0
	for pPow := p; pPow <= n; pPow *= p {
		sum += n / pPow
	}

	return sum
}

func parseMultiplesCode(multiplesCode string) ([]pair, error) {
	powers := strings.Split(multiplesCode, ",")
	ret := make([]pair, len(powers))

	for i := range powers {
		pe := strings.Split(powers[i], "-")
		if len(pe) != 2 {
			return nil, fmt.Errorf("bad arg")
		}

		p, err := strconv.Atoi(pe[0])
		if err != nil {
			return nil, err
		}

		e, err := strconv.Atoi(pe[1])
		if err != nil {
			return nil, err
		}

		ret[i] = pair{p: p, e: e}
	}

	slices.SortFunc(ret, func(x, y pair) int {
		a, b := x.p, y.p

		if a < b {
			return 1
		} else if a > b {
			return -1
		}

		return 1
	})

	return ret, nil
}

/*
	https://en.wikipedia.org/wiki/Kummer%27s_theorem#Multinomial_coefficient_generalization

	valuationPadic(p, n, m1, ..., mk) = 1/(p-1) * (sum[i:1..k]digitSum(p, mi) - digitSum(p, n))

	Legendre p-adic valuation (https://en.wikipedia.org/wiki/P-adic_valuation appied on n factorial) turns out to be much faster
*/
