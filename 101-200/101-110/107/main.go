package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 107; Minimal Network
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
		limit = 1
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	var fileName string
	if limit == 2 {
		fileName = "example.txt"
	} else {
		fileName = "0107_network.txt"
	}

	var rowStrings []string
	if rowStrings, err = projecteuler.FileToStrings(fileName); err != nil {
		return
	}

	matrixDim := len(rowStrings)
	networkMatrix := make([][]int, matrixDim)
	for i, r := range rowStrings {
		networkMatrix[i] = make([]int, matrixDim)
		inputRow := strings.Split(r, ",")

		for j, s := range inputRow {
			if s == "-" {
				networkMatrix[i][j] = 0
				continue
			}

			n, err := strconv.Atoi(s)
			if err != nil {
				panic("bad input")
			}

			networkMatrix[i][j] = n
		}
	}

	resInt := maxSaving(networkMatrix)

	result = strconv.Itoa(resInt)
	return
}

func maxSaving(m [][]int) int {
	dim := len(m)
	mst := minimumSpanningTree(m)

	sumOrig := 0
	sumMST := 0
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			sumOrig += m[i][j]
			sumMST += mst[i][j]
		}
	}

	return (sumOrig - sumMST) / 2
}

// Prim's algorithm
func minimumSpanningTree(m [][]int) [][]int {
	dim := len(m)
	cheapestCost := make([]int, dim)
	explored := map[int]struct{}{}
	unexplored := map[int]struct{}{}

	cheapestEdge := make([]int, dim)
	for i := 0; i < dim; i++ {
		cheapestCost[i] = math.MaxInt
		cheapestEdge[i] = -1
		unexplored[i] = struct{}{}
	}

	cheapestEdge[0] = 0
	cheapestCost[0] = 0

	for len(unexplored) > 1 {
		currVertex := findMin(unexplored, cheapestCost)
		delete(unexplored, currVertex)
		explored[currVertex] = struct{}{}

		for i := 0; i < dim; i++ {
			if m[currVertex][i] == 0 {
				continue
			}

			if _, present := unexplored[i]; present {
				if m[currVertex][i] < cheapestCost[i] {
					cheapestCost[i] = m[currVertex][i]
					cheapestEdge[i] = currVertex
				}
			}
		}
	}

	res := make([][]int, dim)
	for i := 0; i < dim; i++ {
		res[i] = make([]int, dim)
	}

	for i := 0; i < dim; i++ {
		if cheapestEdge[i] == -1 {
			panic("empty cheapestEdge")
		}

		res[i][cheapestEdge[i]] = m[i][cheapestEdge[i]]
		res[cheapestEdge[i]][i] = m[i][cheapestEdge[i]]
	}

	return res
}

func findMin(unexplored map[int]struct{}, cheapestCost []int) int {
	ret := -1
	min := math.MaxInt
	for index := range unexplored {
		if cheapestCost[index] < min {
			min = cheapestCost[index]
			ret = index
		}
	}

	return ret
}
