package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 19; Counting Sundays

You are given the following information, but you may prefer to do some research for yourself.

    1 Jan 1900 was a Monday.
    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
    A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
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

	projecteuler.TimedStr(calc, limit)
}

const (
	sunday = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	mMap := buildMonthMap()

	var resultInt, start int
	_, start = sundaysFirstInYear(monday, mMap)

	for i := 0; i < limit; i++ {
		var currResult int
		currResult, start = sundaysFirstInYear(start, mMap)
		resultInt += currResult
	}

	result = strconv.Itoa(resultInt)
	return
}

func buildMonthMap() (mMap map[int]int) {
	mMap = make(map[int]int)
	mMap[0] = 0
	mMap[1] = 31
	mMap[2] = 28
	mMap[3] = 31
	mMap[4] = 30
	mMap[5] = 31
	mMap[6] = 30
	mMap[7] = 31
	mMap[8] = 31
	mMap[9] = 30
	mMap[10] = 31
	mMap[11] = 30
	mMap[12] = 31

	return
}

func addModSeven(x, y int) int {
	x += y
	x %= 7
	return x
}

func sundaysFirstInYear(start int, mMap map[int]int) (result, nextStart int) {
	for i := 0; i < 12; i++ {
		start = addModSeven(start, mMap[i])
		if start == sunday {
			result++
		}
	}

	nextStart = addModSeven(start, mMap[12])
	return
}
