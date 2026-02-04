package main

import (
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 120; Square Remainders
Let r be the remainder when (a - 1)^n + (a + 1)^n is divided by a^2.

For example, if a = 7 and n = 3, then r = 42: 6^3 + 8^3 = 728 mod 49 = 42. And as n varies, so too will r, but for a = 7 it turns out that r_max = 42.

For 3 <= a <= 1000, find sum(r_max).
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	res := 0

	for a := 3; a < 1001; a++ {
		aa := a * a
		maxR := 0

		for k := 0; k < a; k++ {
			currR := 2 * (2*k + 1) * a

			for currR >= aa {
				currR -= aa
			}
			if currR > maxR {
				maxR = currR
			}
		}
		res += maxR
	}

	result = strconv.Itoa(res)
	return
}

/*
	sum[i 0..n]([n over i](-1)^i * a^(n - i) + [n over i]a^(n - i)) = 2 * sum[i 0..n/2]([n over 2i]a^(n - i))

	n even
		n = 2k, 2 * sum[i 0..k]([2k over 2i]a^(2k - 2i))
			n=2: 2 * ([2 over 0]a^2 + [2 over 2]) = 2 * (a^2 + 1), mod a^2 = 2
			n=6: 2 * ([6 over 0]a^6 + [6 over 2]a^4 + [6 over 4]a^2 + [6 over 6]), mod a^2 = 2

			for even n, r = 2.
	n odd
		n = 2k + 1, 2 * sum[i 0..k]([2k+1 over 2i]a^(2k+1 - 2i))
			n=3: 2 * ([3 over 0]a^3 + [3 over 1]a) = 2 * (a^3 + 3a), mod a^2 = 6a
			n=7: 2 * ([7 over 0]a^7 + [7 over 2]a^5 + [7 over 4]a^3 + [7 over 1]a) = 2 * (a^7 + 21a^5 + 35a^3 + 7a), mod a^2 = 14a
			r = sum mod a^2 = 2*n*a

			For 2*n mod a = x, 2*n = a*y + x => 2*n*a = a^2*y + x*a => 2*n*a mod a^2

			max(a)[2*n*a mod a^2] = ?

			2*(2*k + 1)*a mod a^2 2,6,10,14,18,22,26,30,34,38,42,...
				a=3: mod 9  [6,0,3,6,0,3,6,0,3]
				a=4: mod 16 [8,8,8,...]
				a=6: mod 36 [12,0,24,12,0,24]
				a=7: mod 49 [14,42,21,0,28,7,35,14,42]
				a=8:mod 64   [16,48,16,48,16...]
				a=10:mod 100 [20,60,0,40,80,20,60]
				a=11:mod 121 [22,66,110,33,77,0,44,88,11,55,99,22,66]
*/
