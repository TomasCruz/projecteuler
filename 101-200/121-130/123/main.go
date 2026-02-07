package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 123; Prime Square Remainders
Let p_n be the nth prime: 2, 3, 5, 7, 11, ..., and let r be the remainder when (p_n - 1)^n + (p_n + 1)^n is divided by p_n^2.
For example, when n = 3, p_3 = 5, and 4^3 + 6^3 = 280 (mod 25).

The least value of n for which the remainder first exceeds 10^9 is 7037.
Find the least value of n for which the remainder first exceeds 10^10.
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
		limit = 10
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	power := uint64(1)
	for i := 0; i < limit-3; i++ {
		power *= 10
	}
	pLimit := power
	power *= 1000

	var i int
	projecteuler.PrimesEratosthenes(pLimit, func(args ...any) bool {
		power := args[0].(uint64)
		iPtr := args[1].(*int)
		primes := args[2].([]uint64)

		last := len(primes) - 1
		x := 2 * (uint64(last) + 1) * primes[last]

		if x > power {
			ret := last + 1
			if last%2 == 1 {
				ret++
			}
			*iPtr = ret

			return true
		}

		return false
	}, power, &i)

	result = strconv.Itoa(i)
	return
}

/*
	sum[i 0..n]([n over i](-1)^i * p_n^(n - i) + [n over i]p_n^(n - i)) = 2 * sum[i 0..n/2]([n over 2i]p_n^(n - i))

	n even
		n = 2k, 2 * sum[i 0..k]([2k over 2i]p_n^(2k - 2i))
			n=2: 2 * ([2 over 0]p_n^2 + [2 over 2]) = 2 * (p_n^2 + 1), mod p_n^2 = 2
			n=6: 2 * ([6 over 0]p_n^6 + [6 over 2]p_n^4 + [6 over 4]p_n^2 + [6 over 6]), mod p_n^2 = 2

			for even n, r = 2.
	n odd
		n = 2k + 1, 2 * sum[i 0..k]([2k+1 over 2i]p_n^(2k+1 - 2i))
			n=3: 2 * ([3 over 0]p_n^3 + [3 over 1]p_n) = 2 * (p_n^3 + 3p_n), mod p_n^2 = 6p_n
			n=7: 2 * ([7 over 0]p_n^7 + [7 over 2]p_n^5 + [7 over 4]p_n^3 + [7 over 1]p_n) = 2 * (p_n^7 + 21p_n^5 + 35p_n^3 + 7p_n), mod p_n^2 = 14p_n
			r = sum mod p_n^2 = 2*n*p_n

			For 2*n mod p_n = x, 2*n = p_n*y + x => 2*n*p_n = p_n^2*y + x*p_n => 2*n*p_n mod p_n^2

			max(p_n)[2*n*p_n mod p_n^2] = ?

			2*n*p_n mod p_n^2
				n=1, p_1=2:  mod 4     2*1*2   mod 4      0
				n=3, p_3=5:  mod 25    2*3*5   mod 25     5
				n=5, p_5=11: mod 121   2*5*11  mod 25    10
				n=7, p_7=19: mod 361  2*7*19  mod 361   266
				n=9, p_9=29: mod 841  2*9*29  mod 841   522
				n=11,p_11=37:mod 1369 2*11*37 mod 1369  814
				...
				n=63,p_63=307: 2*63*307 mod 307^2     38682
				...
				n=7036,p_7036=71039: 2*7036*71039  999660808
				n=7037,p_7037=71059: 2*7037*71059 1000084366
*/
