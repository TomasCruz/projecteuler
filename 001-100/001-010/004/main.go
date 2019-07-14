package main

import (
	"fmt"
	"sort"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 4
Largest palindrome product

A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) {
	var found int
	var doBreak bool

	palindromes := findPalindromes()
	for i := 0; i < len(palindromes); i++ {
		for j := 100; j < 1000; j++ {
			if palindromes[i]%j == 0 && isThreeDigits(palindromes[i]/j) {
				found = palindromes[i]
				doBreak = true
				break
			}
		}

		if doBreak {
			break
		}
	}

	fmt.Println(found)
}

func isThreeDigits(num int) bool {
	return num > 99 && num < 1000
}

func findPalindromes() (palindromes []int) {
	var a, b, c int

	for a = 1; a < 10; a++ {
		for b = 0; b < 10; b++ {
			for c = 0; c < 10; c++ {
				palindromes = append(palindromes, 100001*a+10010*b+1100*c)
			}
		}
	}

	sort.Slice(palindromes, func(i, j int) bool {
		return palindromes[i] > palindromes[j]
	})

	return
}
