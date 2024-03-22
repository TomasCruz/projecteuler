package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 80; Square Root Digital Expansion
It is well known that if the square root of a natural number is not an integer, then it is irrational.
The decimal expansion of such square roots is infinite without any repeating pattern at all.
The square root of two is 1.41421356237309504880..., and the digital sum of the first one hundred decimal digits is 475.

For the first one hundred natural numbers, find the total of the digital sums of the first one hundred decimal digits
for all the irrational square roots.
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

type sqRoot struct {
	multiplier int
	root       int
}

type decimalRoot struct {
	rt           projecteuler.BigInt
	decimalIndex int
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	// squares
	squares := []int{0, 1}
	for i := 2; ; i++ {
		sq := i * i
		if sq > limit {
			break
		}
		squares = append(squares, sq)
	}

	// factorize non squares
	primes := projecteuler.Primes(limit, nil)
	j := 2
	sqRoots := make([]sqRoot, limit+1)
	for i := 2; i <= limit; i++ {
		if j < len(squares) && squares[j] == i {
			j++
			continue
		}

		var factors map[int]int
		factors, err = projecteuler.Factorize(i, primes)
		if err != nil {
			return
		}

		mulFactors := map[int]int{}
		root := 1
		for k := range factors {
			v := factors[k]
			rest := v % 2
			mulFactors[k] = v / 2
			if rest == 1 {
				root *= k
			}
		}

		mul := int(projecteuler.MultiplyFactors(mulFactors))
		sqRoots[i] = sqRoot{
			multiplier: mul,
			root:       root,
		}
	}

	// expansions
	j = 2
	expansions := map[int]decimalRoot{}
	for i := 2; i <= limit; i++ {
		if j < len(squares) && squares[j] <= i {
			j++
			if squares[j-1] == i {
				continue
			}
		}

		r := sqRoots[i].root
		if _, present := expansions[r]; !present {
			rBigInt := projecteuler.MakeBigIntFromInt(r)
			root, decimalIndex := rBigInt.SquareRoot(101)

			expansions[r] = decimalRoot{
				rt:           root,
				decimalIndex: decimalIndex,
			}
		}
	}

	// total
	total := 0
	for i := 2; i <= limit; i++ {
		m := sqRoots[i].multiplier
		r := sqRoots[i].root

		// skip squares
		if m == 0 && r == 0 {
			continue
		}

		rtClone := expansions[r].rt.Clone()
		rtClone.MultiplyByDigit(byte(m))

		currExpansion := removeDigits(rtClone, 0, rtClone.DigitCount()-100)
		total += currExpansion.DigitSum()
	}

	result = strconv.Itoa(total)
	return
}

func removeDigits(bi *projecteuler.BigInt, leading, trailing int) projecteuler.BigInt {
	digits := bi.Digits()
	digits = digits[:len(digits)-leading]
	digits = digits[trailing:]
	reversePlusZero(digits)
	res := projecteuler.MakeBigInt(string(digits))
	return res
}

func reversePlusZero(x []byte) {
	l := len(x)
	limit := l / 2

	for i := 0; i < limit; i++ {
		x[i], x[l-1-i] = x[l-1-i], x[i]
	}

	for i := 0; i < l; i++ {
		x[i] = x[i] + '0'
	}
}
