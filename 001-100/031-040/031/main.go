package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 31; Coin sums

In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

    1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).

It is possible to make £2 in the following way:

    1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

How many different ways can £2 be made using any number of coins?
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
		limit = 200
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	fv := []int{1, 2, 5, 10, 20, 50, 100, 200}
	limitIndex := -1

	for i := 0; i < 8; i++ {
		if limit == fv[i] {
			limitIndex = i
			break
		}
	}

	if limitIndex == -1 {
		err = fmt.Errorf("Wrong limit value")
		return
	}

	arrays := buildArrays(limitIndex, fv)
	total := numberWays(2*fv[limitIndex], limitIndex, fv, arrays)

	result = strconv.Itoa(total)
	return
}

func buildArrays(limitIndex int, fv []int) (arrays []map[int]int) {
	arrays = make([]map[int]int, limitIndex+1)

	for i := 2; i <= limitIndex; i++ {
		arrays[i] = make(map[int]int)
	}

	for i := 2; i <= limitIndex; i++ {
		singleArray(i, limitIndex, fv, arrays)
	}

	return
}

func singleArray(index, limitIndex int, fv []int, arrays []map[int]int) {
	limit := fv[limitIndex]
	x := 2 * fv[index]
	for x <= limit {
		numberWays(x, index, fv, arrays)
		x += fv[index]
	}

	return
}

func numberWays(x, index int, fv []int, arrays []map[int]int) int {
	if index < 2 {
		if index == 0 {
			return 1
		}
		return x / 2
	}

	if fv[index] >= x {
		if fv[index] == x {
			return 1
		}
		return 0
	}

	indexMap := arrays[index]
	if val, ok := indexMap[x]; ok {
		return val
	}

	sum := 0
	smaller := x - fv[index]
	for i := 0; i <= index; i++ {
		sum += numberWays(smaller, i, fv, arrays)
	}

	indexMap[x] = sum
	return sum
}
