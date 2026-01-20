package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 109; Darts
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
		limit = 100
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	populateMap()

	sum := 0
	for i := 2; i < limit; i++ {
		sum += checkoutCount(i)
	}

	result = strconv.Itoa(sum)
	return
}

var dartThrows = map[string]int{}

func populateMap() {
	dartThrows["S25"] = 25
	dartThrows["D25"] = 50

	for i := 1; i < 21; i++ {
		dartThrows[fmt.Sprintf("S%02d", i)] = i
		dartThrows[fmt.Sprintf("D%02d", i)] = 2 * i
		dartThrows[fmt.Sprintf("T%02d", i)] = 3 * i
	}
}

func checkoutCount(target int) int {
	checkouts := map[string]struct{}{}
	distinctCheckouts := map[string]struct{}{}

	checkoutRec(target, 3, "", checkouts)

	for ch := range checkouts {
		if len(ch) == 12 {
			equivalentCh := fmt.Sprintf(" %s %s %s", ch[5:8], ch[1:4], ch[9:])
			if _, present := distinctCheckouts[equivalentCh]; present {
				continue
			}
		}

		distinctCheckouts[ch] = struct{}{}
	}

	// for s := range distinctCheckouts {
	// 	fmt.Println(s)
	// }
	return len(distinctCheckouts)
}

func checkoutRec(remainder, remainingThrows int, chSoFar string, currCheckouts map[string]struct{}) {
	if remainingThrows == 0 || remainder == 1 {
		return
	}

	for s, num := range dartThrows {
		if remainder < num {
			continue
		}

		if remainingThrows == 1 && s[0] != 'D' {
			continue
		}

		nextChSoFar := fmt.Sprintf("%s %s", chSoFar, s)

		if remainder > num {
			checkoutRec(remainder-num, remainingThrows-1, nextChSoFar, currCheckouts)
		} else if remainder == num && s[0] == 'D' {
			currCheckouts[nextChSoFar] = struct{}{}
		}
	}
}
