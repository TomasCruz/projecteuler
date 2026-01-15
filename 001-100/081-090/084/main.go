package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
		limit = 4
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	steps := 8000
	games := 100000
	ch := make(chan []fld, games)

	sumBoard := newBoard()
	for i := 0; i < games; i++ {
		go func(ch chan<- []fld) {
			board := newBoard()
			boardIndex := 0
			for j := 0; j < steps; j++ {
				boardIndex = board.makeMove(limit, boardIndex, 0)
			}
			ch <- board.fields
		}(ch)
	}

	for i := 0; i < games; i++ {
		f := <-ch
		for j := 0; j < 40; j++ {
			sumBoard.fields[j].hits += f[j].hits
		}
	}

	sort.Slice(sumBoard.fields, func(i, j int) bool {
		return sumBoard.fields[i].hits > sumBoard.fields[j].hits
	})

	// sumBoard.print()
	result = fmt.Sprintf("%02d%02d%02d", sumBoard.fields[0].ord, sumBoard.fields[1].ord, sumBoard.fields[2].ord)
	return
}
