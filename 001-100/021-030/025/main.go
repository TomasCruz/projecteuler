package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 25; 1000-digit Fibonacci number

The Fibonacci sequence is defined by the recurrence relation:

    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.

Hence the first 12 terms will be:

    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144

The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
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

	var biggestIndex int
	powerMatrices, biggestIndex := buildFibPowerIndices(limit)
	firstIndex := biggestIndex / 2
	lastIndex := biggestIndex
	index := binarySearch(limit, powerMatrices, firstIndex, lastIndex)

	result = strconv.Itoa(index)
	return
}

func buildFibPowerIndices(limit int) (powerMatrices map[int]*projecteuler.BigIntMatrix, biggestIndex int) {
	powerMatrices = make(map[int]*projecteuler.BigIntMatrix)
	one, _ := projecteuler.NewBigIntMatrix(2, 2, []int64{1, 1, 1, 0})
	powerMatrices[1] = one
	biggestIndex = 1

	for {
		prev := powerMatrices[biggestIndex]
		biggestIndex *= 2
		curr := projecteuler.MulBigIntMatrix(prev, prev)
		powerMatrices[biggestIndex] = curr

		fn, _ := curr.At(0, 1)
		if len(fn.String()) >= limit {
			break
		}
	}

	return
}

func binarySearch(limit int, powerMatrices map[int]*projecteuler.BigIntMatrix, firstIndex, lastIndex int) int {
	if lastIndex-firstIndex <= 1 {
		return lastIndex
	}

	middleIndex := (firstIndex + lastIndex) / 2
	insertIndex(powerMatrices, middleIndex)

	found, belowLimit, offset := checkMatrix(limit, powerMatrices[middleIndex])
	if found {
		return middleIndex + offset
	}

	if belowLimit {
		return binarySearch(limit, powerMatrices, middleIndex, lastIndex)
	}

	return binarySearch(limit, powerMatrices, firstIndex, middleIndex)
}

func insertIndex(powerMatrices map[int]*projecteuler.BigIntMatrix, middleIndex int) {
	if _, ok := powerMatrices[middleIndex]; ok {
		return
	}

	sl := make([]int, 0)
	powerTwoAdditionTerms(middleIndex, &sl)
	result := powerMatrices[sl[0]]
	for i := 1; i < len(sl); i++ {
		temp := projecteuler.MulBigIntMatrix(result, powerMatrices[sl[i]])
		result = temp.Clone()
	}

	powerMatrices[middleIndex] = result
}

func powerTwoAdditionTerms(middleIndex int, slice *[]int) {
	if middleIndex == 0 {
		return
	}

	biggest := middleIndex
	cnt := 0
	for biggest != 1 {
		biggest >>= 1
		cnt++
	}

	for cnt != 0 {
		biggest <<= 1
		cnt--
	}

	*slice = append(*slice, biggest)
	powerTwoAdditionTerms(middleIndex-biggest, slice)
}

func checkMatrix(limit int, matrix *projecteuler.BigIntMatrix) (found, belowLimit bool, offset int) {
	fn, _ := matrix.At(0, 1)
	fPrev, _ := matrix.At(1, 1)
	fNext, _ := matrix.At(0, 0)

	digitsN := len(fn.String())
	digitsPrev := len(fPrev.String())
	digitsNext := len(fNext.String())

	if digitsPrev >= limit {
		return
	}

	if digitsN == limit {
		found = true
		offset = 0
		return
	}

	if digitsNext == limit {
		if digitsN < limit {
			found = true
			offset = 1
			return
		}
	} else {
		// digitsNext < limit
		belowLimit = true
		return
	}

	return
}
