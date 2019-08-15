package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 36; Double-base palindromes

The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
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
		limit = 1000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	sum := 0
	for i := 1; i < limit; i += 2 {
		base10 := extractDigits(i, 10)
		base2 := extractDigits(i, 2)

		if palindrome(base10) && palindrome(base2) {
			sum += i
		}
	}

	result = strconv.Itoa(sum)
	return
}

func extractDigits(x, base int) (digits []int) {
	for x > 0 {
		digits = append(digits, x%base)
		x /= base
	}

	return
}

func palindrome(digits []int) bool {
	length := len(digits)
	for i := 0; i < length/2; i++ {
		if digits[i] != digits[length-1-i] {
			return false
		}
	}

	return true
}
