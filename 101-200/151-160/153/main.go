package main

import (
	"fmt"
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

	coprimeSum := findCoprimeSum(limit)
	equalSum := findEqualSum(limit)
	realSum := findRealSum(limit)
	sum := coprimeSum + equalSum + realSum

	result = strconv.FormatInt(sum, 10)
	return
}

type pair struct {
	a, b, norm, sum int
}

func newPair(a, b int) pair {
	return pair{
		a:    a,
		b:    b,
		norm: a*a + b*b,
		sum:  a + b,
	}
}

func (p pair) String() string {
	return fmt.Sprintf("(%d, %d)\t%d, %d", p.a, p.b, p.norm, p.sum)
}

func findCoprimeSum(n int) int64 {
	ret := int64(0)

	for b := 2; b < n; b++ {
		bs := projecteuler.NewBitset(b+1, 64)
		for a := 2; a <= b/2; a++ {
			if b%a == 0 {
				for i := 1; ; i++ {
					mul := a * i
					if mul > b {
						break
					}

					bs.Set(mul, true)
				}
			}
		}

		for a := 1; a < b; a++ {
			if !bs.Get(a) {
				p := newPair(a, b)
				if p.norm > n {
					break
				}

				sum := int64(0)
				for k := 1; ; k++ {
					if k*p.norm > n {
						break
					}
					sum += (int64(n) / (int64(k) * int64(p.norm))) * 2 * int64(k) * int64(p.sum)
				}

				ret += sum
			}
		}
	}

	return ret
}

func findEqualSum(n int) int64 {
	ret := int64(0)

	for k := 1; ; k++ {
		if 2*k > n {
			break
		}
		ret += (int64(n) / (int64(k) * 2)) * 2 * int64(k)
	}

	return ret
}

func findRealSum(n int) int64 {
	ret := int64(0)

	for i := 1; i <= n; i++ {
		ret += (int64(n) / int64(i)) * int64(i)
	}

	return ret
}

/*
	Wanted sum of sum of divisors is sum of sum of complex Gaussian integers divisors plus sum of sum of real divisors,
	which is sum of P[1..r](pi^a[i+1]-1)/(pi-1)) == sum[1..n](i*floor(n/i))

	https://math.stackexchange.com/questions/1504639/how-to-simplify-a-sum-of-complex-divisors

	n = (a+bi)*(c+di) = (ac-bd) + (bc+ad)i
	norm of divisor has to divide n

	Every true complex divisor with a positive real part can be expressed as k*(a+bi), gcd(a,b)=1 and a,b,k positive. Let norm be c = a^2 + b^2.
	For a != b:
		For every k, there is floor(n/kc) multiples of k*(a+bi), k*(a-bi), k*(b-ai) and k*(b+ai), a+bi multiplied by an identity.
		Sum of real parts of these is 2*k*(a+b). It is therefore enough to consider only a < b.
	For a = b:
		a=b=1, c=2. Sum is not 4k but 2k, because 1+i multiplied by an identity is 1+i,-1-i(out),-1+i(out) and 1-i so sum is sum[1..n/2](floor(n/2k)*2k)

	Overall sum is sum[a,b | c <= n && gcd(a,b)=1 && a < b](sum[1..n](floor(n/kc) * 2*k*(a+b)) + sum[1..n/2](floor(n/2k)*2k) + sum[1..n](i*floor(n/i))
*/
