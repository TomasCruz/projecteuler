package main

import (
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 89; Roman Numerals


For a number written in Roman numerals to be considered valid there are basic rules which must be followed.
Even though the rules allow some numbers to be expressed in more than one way there is always a "best" way of writing a particular number.
For example, it would appear that there are at least six ways of writing the number sixteen:

IIIIIIIIIIIIIIII
VIIIIIIIIIII
VVIIIIII
XIIIIII
VVVI
XVI

However, according to the rules only XIIIIII and XVI are valid, and the last example is considered to be the most efficient,
as it uses the least number of numerals.
The 11K text file, roman.txt (right click and 'Save Link/Target As...'), contains one thousand numbers written in valid, but not necessarily minimal,
Roman numerals; see About... Roman Numerals for the definitive rules for this problem.

Find the number of characters saved by writing each of these in their minimal form.
Note: You can assume that all the Roman numerals in the file contain no more than four consecutive identical units.
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	var numerals []string
	if numerals, err = projecteuler.FileToStrings("0089_roman.txt"); err != nil {
		return
	}

	res := 0
	for _, s := range numerals {
		res += len(s) - len(minForm(s))
	}
	result = strconv.Itoa(res)

	return
}

const (
	IntI int = 1
	IntV int = 5
	IntX int = 10
	IntL int = 50
	IntC int = 100
	IntD int = 500
	IntM int = 1000

	RuneI rune = 'I'
	RuneV rune = 'V'
	RuneX rune = 'X'
	RuneL rune = 'L'
	RuneC rune = 'C'
	RuneD rune = 'D'
	RuneM rune = 'M'
)

func minForm(numeral string) string {
	value := parseNumeral(numeral)
	return makeMinForm(value)
}

func parseNumeral(numeral string) int {
	value := 0

	for i := 0; i < len(numeral); i++ {
		switch rune(numeral[i]) {
		case RuneI:
			if i < len(numeral)-1 && (rune(numeral[i+1]) == RuneV || rune(numeral[i+1]) == RuneX) {
				value -= 1
			} else {
				value += 1
			}
		case RuneV:
			value += 5
		case RuneX:
			if i < len(numeral)-1 && (rune(numeral[i+1]) == RuneL || rune(numeral[i+1]) == RuneC) {
				value -= 10
			} else {
				value += 10
			}
		case RuneL:
			value += 50
		case RuneC:
			if i < len(numeral)-1 && (rune(numeral[i+1]) == RuneD || rune(numeral[i+1]) == RuneM) {
				value -= 100
			} else {
				value += 100
			}
		case RuneD:
			value += 500
		case RuneM:
			value += 1000
		}
	}

	return value
}

func makeMinForm(value int) string {
	var sb strings.Builder

	// 1000s
	for value > IntM {
		sb.WriteRune(RuneM)
		value -= IntM
	}

	// 100s
	if value >= 900 {
		sb.WriteRune(RuneC)
		sb.WriteRune(RuneM)
		value -= 900
	} else if value >= IntD {
		sb.WriteRune(RuneD)
		value -= IntD
	} else if value >= 400 {
		sb.WriteRune(RuneC)
		sb.WriteRune(RuneD)
		value -= 400
	}

	for value >= IntC {
		sb.WriteRune(RuneC)
		value -= IntC
	}

	// 10s
	if value >= 90 {
		sb.WriteRune(RuneX)
		sb.WriteRune(RuneC)
		value -= 90
	} else if value >= IntL {
		sb.WriteRune(RuneL)
		value -= IntL
	} else if value >= 40 {
		sb.WriteRune(RuneX)
		sb.WriteRune(RuneL)
		value -= 40
	}

	for value >= IntX {
		sb.WriteRune(RuneX)
		value -= IntX
	}

	// 1s
	if value >= 9 {
		sb.WriteRune(RuneI)
		sb.WriteRune(RuneX)
		value -= 9
	} else if value >= IntV {
		sb.WriteRune(RuneV)
		value -= IntV
	} else if value == 4 {
		sb.WriteRune(RuneI)
		sb.WriteRune(RuneV)
		value -= 4
	}

	for value >= IntI {
		sb.WriteRune(RuneI)
		value -= IntI
	}

	return sb.String()
}
