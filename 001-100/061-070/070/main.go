package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 70;
Euler's totient function, phi(n) [sometimes called the phi function], is used to determine the number of positive numbers
less than or equal to n which are relatively prime to n. For example, as 1, 2, 4, 5, 7 and 8, are all less than or equal
to nine and relatively prime to nine, phi(9) = 6.
The number 1 is considered to be relatively prime to every positive number, so phi(1) = 1.

Interestingly, phi(87109) = 79180, and it can be seen that 87109 is a permutation of 79180.

Find the value of n, 1 < n < 10^7, for which phi(n) is a permutation of n and the ratio n/phi(n) produces a minimum.
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
		limit = 10000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	m := map[int]map[int]struct{}{} // smallest perm -> set of perms

	checked := map[int]struct{}{}
	for n := 2; n < limit; n++ {
		if _, present := checked[n]; present {
			continue
		}

		smallest, digitPerms := permMap(n)
		if _, present := m[smallest]; !present {
			m[smallest] = digitPerms
			for p := range digitPerms {
				checked[p] = struct{}{}
			}
		}
	}

	minX := 0
	minRatio := math.MaxFloat64
	primes := projecteuler.Primes(limit, nil)
	for smallest, perms := range m {
		if len(perms) == 1 {
			continue
		}

		for x := range perms {
			if x == smallest {
				continue
			}

			t := projecteuler.Totient(x, primes)
			if _, present := perms[t]; present {
				ratio := float64(x) / float64(t)
				if ratio < minRatio {
					minX = x
					minRatio = ratio
				}
			}
		}
	}

	result = strconv.Itoa(minX)
	return
}

type digitalNumber struct {
	x      int
	digits string
}

func newDigitalNumber(pdn projecteuler.DigitalNumber) digitalNumber {
	digs := pdn.Digits()
	l := len(digs)
	digits := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		digits = append(digits, digs[i]+'0')
	}

	// reverse it
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digitalNumber{
		x:      pdn.X(),
		digits: string(digits),
	}
}

func newDigitalNumberFromString(s string) digitalNumber {
	x := 0
	pow := 1
	l := len(s)
	for i := l; i > 0; i-- {
		x += pow * int(s[i-1]-'0')
		pow *= 10
	}

	return digitalNumber{
		x:      x,
		digits: s,
	}
}

func permMap(x int) (smallest int, digitPerms map[int]struct{}) {
	peDigitalX := projecteuler.NewDigitalNumber(x)
	digitalX := newDigitalNumber(peDigitalX)
	l := len(digitalX.digits)

	perms := projecteuler.Permutations(byte(l), nil)
	smallest = math.MaxInt
	digitPerms = map[int]struct{}{}
	for _, p := range perms {
		if digitalX.digits[p[0]] == '0' {
			continue
		}

		currDigits := make([]byte, l)
		for i := 0; i < l; i++ {
			currDigits[i] = digitalX.digits[p[i]]
		}

		dn := newDigitalNumberFromString(string(currDigits))
		if dn.x < smallest {
			smallest = dn.x
		}
		digitPerms[dn.x] = struct{}{}
	}

	return
}

func phi(n int, primes []int) int {
	// Initialize result as n
	result := float64(n)

	// Consider all prime factors of n and for every prime factor p, multiply result with (1 - 1/p)
	for i := 0; primes[i]*primes[i] <= n; i++ {
		// Check if it is a prime factor.
		if n%primes[i] == 0 {
			// update n and result
			for n%primes[i] == 0 {
				n /= primes[i]
			}
			result *= (1.0 - (1.0 / float64(primes[i])))
		}
	}

	// If n has a prime factor greater than sqrt(n), there can be at-most one such prime factor
	if n > 1 {
		result -= result / float64(n)
	}

	//Since in the set {1,2,....,n-1}, all numbers are relatively prime with n
	//if n is a prime number

	return int(result)
}
