package main

import (
	"fmt"
	"math"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 151; A Preference for A5
A printing shop runs 16 batches (jobs) every week and each batch requires a sheet of special colour-proofing paper of size A5.

Every Monday morning, the supervisor opens a new envelope, containing a large sheet of the special paper with size A1.

The supervisor proceeds to cut it in half, thus getting two sheets of size A2. Then one of the sheets is cut in half to get two sheets of size A3
and so on until an A5-size sheet is obtained, which is needed for the first batch of the week.

All the unused sheets are placed back in the envelope.

At the beginning of each subsequent batch, the supervisor takes from the envelope one sheet of paper at random. If it is of size A5, then it is used.
If it is larger, then the 'cut-in-half' procedure is repeated until an A5-size sheet is obtained, and any remaining sheets are always placed back in the envelope.

Excluding the first and last batch of the week, find the expected number of times (during each week)
that the supervisor finds a single sheet of paper in the envelope.

Give your answer rounded to six decimal places using the format x.xxxxxx .
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	sheets := [4]int{1, 1, 1, 1}
	sum := step(sheets, 1, 0)
	sum = math.Round(1000000*sum) / 1000000.0

	result = fmt.Sprintf("%.6f", sum)
	return
}

func step(sheets [4]int, p, sum float64) float64 {
	numSheets := 0
	for i := range 4 {
		numSheets += sheets[i]
	}

	if numSheets == 1 {
		if sheets[3] == 1 {
			return sum
		}

		sum += p
	}

	for i := range 4 {
		if sheets[i] == 0 {
			continue
		}

		currP := p * float64(sheets[i]) / float64(numSheets)
		nextSheets := cut(sheets, i)
		sum = step(nextSheets, currP, sum)
	}

	return sum
}

func cut(sheets [4]int, index int) [4]int {
	res := [4]int{}
	copy(res[:], sheets[:])

	res[index]--
	for i := index + 1; i < 4; i++ {
		res[i]++
	}

	return res
}
