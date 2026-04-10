package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 153; Investigating Gaussian Integers
As we all know the equation x^2=-1 has no solutions for real x.

If we however introduce the imaginary number i this equation has two solutions: x=i and x=-i.

If we go a step further the equation (x-3)^2=-4 has two complex solutions: x=3+2i and x=3-2i.
x=3+2i and x=3-2i are called each others' complex conjugate.

Numbers of the form a+bi are called complex numbers.

In general a+bi and a-bi are each other's complex conjugate.

A Gaussian Integer is a complex number a+bi such that both a and b are integers.

The regular integers are also Gaussian integers (with b=0).

To distinguish them from Gaussian integers with b \ne 0 we call such integers "rational integers."

A Gaussian integer a+bi is called a divisor of a rational integer n if the result n / (a + bi) is also a Gaussian integer.

If for example we divide 5 by 1+2i we can simplify 5/(1 + 2i) in the following manner:

Multiply numerator and denominator by the complex conjugate of 1+2i: 1-2i.

The result is 5/(1 + 2i) = 5/(1 + 2i)*(1 - 2i/(1 - 2i) = 5(1 - 2i)/(1 - (2i)^2) = 5(1 - 2i)/(1 - (-4)) = 5(1 - 2i)/(5) = 1 - 2i.

So 1+2i is a divisor of 5.

Note that 1+i is not a divisor of 5 because 5/(1 + i) = 5/2 - 5/2*i.

Note also that if the Gaussian Integer (a+bi) is a divisor of a rational integer n, then its complex conjugate (a-bi) is also a divisor of n.

In fact, 5 has six divisors such that the real part is positive: {1, 1 + 2i, 1 - 2i, 2 + i, 2 - i, 5}.

The following is a table of all of the divisors for the first five positive rational integers:
n	 Gaussian integer divisors with positive real part		Sum s(n) of these divisors
1		1													1
2		1, 1+i, 1-i, 2										5
3		1, 3												4
4		1, 1+i, 1-i, 2, 2+2i, 2-2i, 4						13
5		1, 1+2i, 1-2i, 2+i, 2-i, 5							12

For divisors with positive real parts, then, we have: sum {1,5} (s(n)) = 35.

sum {1,10^5} (s(n)) = 17924657155.

What is sum {1,10^8} (s(n))?
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
		limit = 100000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	result = strconv.Itoa(limit)
	return
}

func findK(p int64) int64 {
	for n := int64(2); n < p; n++ {
		// n^(p-1)/2 mod p == -1
		pow := n % p
		exp := (p - 1) / 2
		for i := 1; i < int(exp); i++ {
			pow *= n
			pow %= p
		}

		if pow == -1 {
			return exp / 2
		}
	}

	return 0
}

// extended Euclidean algorithm
func extendedEuclidean(a, n int) int {
	return 0
}

/*
	a,c > 0, sign b and sign d must be different, otherwise ad + bc != 0
	n / (a + bi) = c + di => n = (a + bi) * (c + di) = (ac - bd) + i*(ad + bc)
	n / (a - bi) = c - di => n = (a - bi) * (c - di) = (ac - bd) - i*(ad + bc)
	n / (a + bi) = c + di => n / (a - bi) = c - di

	ad + bc = 0, ac - bd = n

	e.g. 4/(1-i) = 2+2i, a = 1, c = 2, b = -1, d = 2

	Since sign b and sign d must be different, -bd is positive => both ac and -bd are positive => all of a,b,c,d <= n
	For complex Gaussian integers, a,b,c,d < n because bd != 0
	For real Gaussian integers, (b, d = 0), a and c are divisors of n.

	Wanted sum of divisors is sum of real divisors P[1..r](pi^a[i+1]-1)/(pi-1)) plus sum of complex Gaussian integers divisors.

	Without loss of generality, b is positive and d is negative. For all symbols below representing positive integers,
	n / (a+bi) = c-di; bc - ad = 0, ac + bd = n
		=> bc = ad, acd + bd^2 = nd, bc^2 + bd^2 = nd, b*(c^2 + d^2) = nd
	Same way a*(c^2 + d^2) = nc, c*(a^2 + b^2) = na, d*(a^2 + b^2) = nb

	https://stackoverflow.com/questions/2269810/whats-a-nice-method-to-factor-gaussian-integers
	https://en.wikipedia.org/wiki/Euclidean_algorithm#Gaussian_integers
	https://mathworld.wolfram.com/DivisorFunction.html (end section, related to Gaussian integers)
	https://en.wikipedia.org/wiki/Modular_multiplicative_inverse#Extended_Euclidean_algorithm
*/
