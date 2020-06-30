package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 55; Lychrel numbers

If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.
Not all numbers produce palindromes so quickly. For example,

349 + 943 = 1292,
1292 + 2921 = 4213
4213 + 3124 = 7337

That is, 349 took three iterations to arrive at a palindrome. Although no one has proved it yet,
it is thought that some numbers, like 196, never produce a palindrome. A number that never forms a palindrome
through the reverse and add process is called a Lychrel number. Due to the theoretical nature of these numbers,
and for the purpose of this problem, we shall assume that a number is Lychrel until proven otherwise.

In addition you are given that for every number below ten-thousand, it will either
(i) become a palindrome in less than fifty iterations, or,
(ii) no one, with all the computing power that exists, has managed so far to map it to a palindrome.

In fact, 10677 is the first number to be shown to require over fifty iterations before producing a palindrome:
4668731596684224866951378664 (53 iterations, 28-digits).
Surprisingly, there are palindromic numbers that are themselves Lychrel numbers; the first example is 4994.

How many Lychrel numbers are there below ten-thousand?
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
		limit = 10000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	result64 := int64(0)
	for i := 1; i < limit; i++ {
		if isLychrel(i) {
			result64++
		}
	}
	result = strconv.FormatInt(result64, 10)

	return
}

func isLychrel(x int) bool {
	num := projecteuler.MakeBigIntFromInt(x)
	arg := num.Clone()

	for i := 1; i < 50; i++ {
		argRev := arg.Clone()
		argRev.ReverseDigits()
		arg.AddTo(argRev)
		if arg.IsPalindrome() {
			return false
		}
	}

	return true
}
