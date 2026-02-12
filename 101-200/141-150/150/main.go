package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 150; Sub-triangle Sums
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
		limit = 2
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	var seq []int
	if limit == 1 {
		seq = []int{15, -14, -7, 20, -13, -5, -3, 8, 23, -26, 1, -4, -5, -18, 5, -16, 31, 2, 9, 28, 3}
	} else {
		seq = linearCongruentialGenerator()
	}

	m := buildMatrix(seq)
	n := len(m)

	min := m[0][0]
	for i := n; i > 1; i-- {
		subMatrix := calcSubTriangleMatrix(m[:i])
		currMin := getMin(subMatrix)
		if currMin < min {
			min = currMin
		}
	}

	result = strconv.Itoa(min)
	return
}

func linearCongruentialGenerator() []int {
	two19 := 1
	for i := 0; i < 19; i++ {
		two19 *= 2
	}

	limit := 500500
	seq := make([]int, 0, limit)

	t := 0
	for k := 0; k < limit; k++ {
		t = (615949*t + 797807) % (2 * two19)
		seq = append(seq, t-two19)
	}

	return seq
}

func buildMatrix(seq []int) [][]int {
	n := len(seq) * 2 // x^2+x-a=0, [-1+-sqrt(1+4a)]/2
	sqrt := math.Sqrt(float64(1 + 4*n))
	dim := int((-1.0 + sqrt) / 2.0)

	m := make([][]int, dim)
	mCounter := 0
	for i := 0; i < dim; i++ {
		m[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			m[i][j] = seq[mCounter]
			mCounter++
		}
	}

	return m
}

func calcSubTriangleMatrix(m [][]int) [][]int {
	dim := len(m)
	subM := make([][]int, dim)

	subM[dim-2] = make([]int, dim-1)
	subM[dim-1] = make([]int, dim)
	subM[dim-1][dim-1] = m[dim-1][dim-1]
	for i := dim - 1; i > 0; i-- {
		subM[dim-1][i-1] = m[dim-1][i-1]
		subM[dim-2][i-1] = m[dim-2][i-1] + m[dim-1][i-1] + m[dim-1][i]
	}

	for i := dim - 3; i >= 0; i-- {
		subM[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			subM[i][j] = m[i][j] + subM[i+1][j] + subM[i+1][j+1] - subM[i+2][j+1]
		}
	}

	return subM
}

func getMin(m [][]int) int {
	dim := len(m)
	min := math.MaxInt

	for i := 0; i < dim; i++ {
		for j := 0; j <= i; j++ {
			if m[i][j] < min {
				min = m[i][j]
			}
		}
	}

	return min
}
