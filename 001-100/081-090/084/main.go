package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 84; Monopoly Odds
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
		limit = 6 //4
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	probs := buildProbabilitiesSingle(limit)
	probMatrix := make([][][]fld, 3)
	probMatrix[0] = make([][]fld, 40)
	probMatrix[1] = make([][]fld, 40)
	probMatrix[2] = make([][]fld, 40)

	// i := 0
	// board := calcProbs(limit, i, 2, probs)
	// printBoard(board)
	// probSum(i, board)

	for j := 3; j > 0; j-- {
		for i := 0; i < 40; i++ {
			probMatrix[j-1][i] = calcProbs(limit, i, j-1, probs, probMatrix)
			// printBoard(board)
			probSum(j-1, i, probMatrix[j-1][i])
		}
	}

	return
}

func probSum(rank, index int, board []fld) {
	p := frac{num: 0, denom: 1}
	for i := 0; i < 40; i++ {
		p = p.add(board[i].prob)
	}

	fmt.Printf("%d %d: %s %6f\n", rank, index, p.toString(), float64(p.num)/float64(p.denom))
}
