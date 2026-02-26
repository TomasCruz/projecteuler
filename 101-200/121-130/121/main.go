package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 121; Disc Game Prize Fund
A bag contains one red disc and one blue disc. In a game of chance a player takes a disc at random and its colour is noted.
After each turn the disc is returned to the bag, an extra red disc is added, and another disc is taken at random.

The player pays £1 to play and wins if they have taken more blue discs than red discs at the end of the game.

If the game is played for four turns, the probability of a player winning is exactly 11/120, and so the maximum prize fund
the banker should allocate for winning in this game would be £10 before they would expect to incur a loss.
Note that any payout will be a whole number of pounds and also includes the original £1 paid to play the game,
so in the example given the player actually wins £9.

Find the maximum prize fund that should be allocated to a single game in which fifteen turns are played.
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
		limit = 15
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	totalSum := 0
	minWins := limit/2 + 1
	for i := minWins; i <= limit; i++ {
		currSum := 0
		projecteuler.Combinations(byte(limit), byte(i), func(args ...interface{}) bool {
			sumPtr := args[0].(*int)
			combo := args[1].([]byte)

			curComboValue := 1
			for i := range combo {
				curr := i + 1
				if combo[i] == 1 {
					curr = 1
				}

				curComboValue *= curr
			}
			*sumPtr += curComboValue

			return false
		}, &currSum)

		totalSum += currSum
	}

	prod := 1
	for i := 2; i < limit+2; i++ {
		prod *= i
	}
	res := prod / totalSum

	result = strconv.Itoa(res)
	return
}
