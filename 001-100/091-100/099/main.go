package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 99; Largest Exponential
Comparing two numbers written in index form like 2^11 and 3^7 is not difficult, as any calculator would confirm that 2^11 = 2048 < 3^7 = 2187.
However, confirming that 632382^518061 > 519432^525806 would be much more difficult, as both numbers contain over three million digits.
Using base_exp.txt, a 22K text file containing one thousand lines with a base/exponent pair on each line,
determine which line number has the greatest numerical value.
NOTE: The first two lines in the file represent the numbers in the example given above.
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
		limit = 1000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	var textNumbers []string
	if textNumbers, err = projecteuler.FileToStrings("0099_base_exp.txt"); err != nil {
		fmt.Println(err)
		return
	}

	// 632382^518061 > 519432^525806
	// 518061 * ln(632382) > 525806 * ln(519432)
	// 6919869 > 6919865, yup
	// math.Log()

	greatestIndex := 0
	greatest := float64(1)
	curr := float64(1)
	for i := 1; i <= limit; i++ {
		str := strings.Split(textNumbers[i-1], ",")
		base, _ := strconv.ParseFloat(str[0], 64)
		exp, _ := strconv.ParseFloat(str[1], 64)
		curr = math.Log(base) * exp
		// curr.Exp(big.NewInt(base), big.NewInt(exp), nil)
		if curr > greatest {
			greatest = curr
			greatestIndex = i
		}
	}

	result = strconv.Itoa(greatestIndex)
	return
}
