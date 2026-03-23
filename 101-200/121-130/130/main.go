package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 130; Composites with Prime Repunit Property
A number consisting entirely of ones is called a repunit. We shall define R(k) to be a repunit of length k; for example, R(6) = 111111.

Given that n is a positive integer and gcd(n, 10) = 1, it can be shown that there always exists a value, k,
for which R(k) is divisible by n, and let A(n) be the least such value of k; for example, A(7) = 6 and A(41) = 5.

You are given that for all primes, p > 5, that p - 1 is divisible by A(p). For example, when p = 41, A(41) = 5, and 40 is divisible by 5.

However, there are rare composite values for which this is also true; the first five examples being 91, 259, 451, 481, and 703.

Find the sum of the first twenty-five composite values of n for which gcd(n, 10) = 1 and n - 1 is divisible by A(n).
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
		limit = 25
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	_, primeSet := projecteuler.PrimeSet(1000000)

	n := 0
	base := 0
	mods := []int{1, 3, 7, 9}
	count := 0
	sum := 0

outter:
	for {
		for _, j := range mods {
			n = base + j
			if n == 1 {
				continue
			}

			if _, present := primeSet[n]; present {
				continue
			}

			a := capitalA(n)
			if (n-1)%a == 0 {
				count++
				sum += n
				if count == limit {
					break outter
				}
			}
		}

		base += 10
	}

	result = strconv.Itoa(sum)
	return
}

func capitalA(n int) int {
	pow10 := 10
	k := 1

	for {
		pow10 = pow10 % (9 * n)
		if pow10 == 1 {
			break
		}

		pow10 *= 10
		k++
	}

	return k
}

/*
	User ThePower
	R(a)=111....111 = 1/9*(999...999) = (10^a-1)/9
       a times           a times
	so change the a to solve the equation:

	10^a mod 9*n = 1

	The A(n) is very simple to code :)
	VBA

	Function GetMinN(n)
	ten_power = 10
	k = 1
	While ten_power - 9 * n * Int(ten_power / (9 * n)) <> 1
	k = k + 1
	ten_power = 10 * ten_power
	ten_power = ten_power - 9 * n * Int(ten_power / (9 * n))
	Wend
	GetMinN = k
	End Function
*/
