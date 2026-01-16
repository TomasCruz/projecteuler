package main

import (
	"fmt"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 101; Optimum Polynomial
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	// koefs := []int64{0, 0, 0, 1}
	koefs := []int64{1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1}
	checkLimit := 15
	seq := pgf(koefs, checkLimit)
	// printSeq(seq[1:])
	// fmt.Println()

	ops := getOPs(seq, len(koefs)-1)
	fit := make([]int64, len(koefs)-1)
	for i := range ops {
		opGenerated := pgf(ops[i], checkLimit)
		var j int
		for j = 0; j < len(seq) && seq[j] == opGenerated[j]; j++ {
		}
		if j == len(seq) {
			panic("all equal!")
		}
		fit[i] = opGenerated[j]
	}

	fitSum := int64(0)
	for i := range fit {
		fitSum += fit[i]
	}

	result = strconv.FormatInt(fitSum, 10)
	return
}

func pgf(koefs []int64, terms int) []int64 {
	ret := make([]int64, terms+1)

	for i := range koefs {
		ret[1] += koefs[i]
	}

	for i := 2; i <= terms; i++ {
		for j := range koefs {
			ret[i] += koefs[j] * pow(int64(i), j)
		}
	}

	return ret
}

func printSeq(seq []int64) {
	fmt.Printf("%d", seq[0])

	for i := 1; i < len(seq); i++ {
		fmt.Printf(", %d", seq[i])
	}

	fmt.Println()
}

func pow(n int64, p int) int64 {
	res := int64(1)

	for i := 0; i < p; i++ {
		res *= n
	}

	return res
}

func getOPs(seq []int64, l int) [][]int64 {
	ret := make([][]int64, l)
	ret[0] = []int64{seq[1]}

	for i := 1; i < l; i++ {
		ret[i] = make([]int64, i+1)
		// fmt.Println(i, ":")
		gaussJordanElimination(seq[:i+2], ret[i])
		// printSeq(ret[i])

		/*
			k0 = seq[0]
			-----
			k0 + k1*1 = seq[0]
			k0 + k1*2 = seq[1]
			-----
			k0 + k1*1+k2*1 = seq[0]
			k0 + k1*2 + k2*2^2 = seq[1]
			k0 + k1*3 + k2*3^2 = seq[2]

			| 1 1 1 |   | k0 |   | 1 |
			| 1 2 4 | * | k1 | = | 8 |
			| 1 3 9 |   | k2 |   | 27|
		*/
	}

	return ret
}

func gaussJordanElimination(seq, ret []int64) {
	l := len(seq) - 1
	m := make([][]int64, l)

	for i := 0; i < l; i++ {
		m[i] = make([]int64, l+1)
		for j := 0; j < l; j++ {
			m[i][j] = pow(int64(i+1), j)
		}
		m[i][l] = seq[i+1]
	}
	// projecteuler.PrintMatrix(m)

	for i := 0; i < l; i++ {
		divRow(m[i], m[i][i])
		for j := i + 1; j < l; j++ {
			addMultipliedRow(m[j], m[i], -m[j][i])
		}
		for j := i; j > 0; j-- {
			addMultipliedRow(m[j-1], m[i], -m[j-1][i])
		}
	}

	for i := 0; i < l; i++ {
		ret[i] = m[i][l]
	}
}

func divRow(row []int64, d int64) {
	if d == 1 {
		return
	}

	for i := range row {
		if row[i]%d != 0 {
			panic("argh")
		}

		row[i] /= d
	}
}

func addMultipliedRow(f, s []int64, mul int64) {
	for i := range f {
		f[i] += s[i] * mul
	}
}
