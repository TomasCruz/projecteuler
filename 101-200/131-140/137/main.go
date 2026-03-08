package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 137; Fibonacci Golden Nuggets
Consider the infinite polynomial series A_F(x) = x*F1 + x^2*F2 + x^3*F3 + ...,
where Fk is the kth term in the Fibonacci sequence: 1, 1, 2, 3, 5, 8, ...; that is, Fk = Fk-1 + Fk-2, F1 = 1 and F2 = 1.

For this problem we shall be interested in values of x for which A_F(x) is a positive integer.

Surprisingly
	A_F(1/2) = (1/2) * 1 + (1/2)^2 * 1 + (1/2)^3* 2 + (1/2)^4* 3 + (1/2)^5* 5 + ... =
		= 1/2 + 1/4 + 2/8 + 3/16 + 5/32 + ... = 2

The corresponding values of x for the first five natural numbers are shown below.

x	A_F(x)</th>
-------------------
sqrt(2)-1		1
1/2				2
(sqrt(13)-2)/3	3
(sqrt(89)-5)/8	4
(sqrt(34)-3)/5	5

We shall call A_F(x) a golden nugget if x is rational, because they become increasingly rarer; for example, the 10th golden nugget is 74049690.

Find the 15th golden nugget.
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
		limit = 15
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	nug := make([]int64, limit+1)

	a, b := int64(1), int64(1)
	i := 1
	for i <= limit {
		a, b = (3*a+5*b)/2, (3*b+a)/2
		if a%5 != 1 {
			continue
		}

		n := (a - 1) / 5
		nug[i] = n
		i++
	}

	result = strconv.FormatInt(nug[limit], 10)
	return
}

/*
	Approach with Pythagorean triplets works decently until eating up RAM.
	Brutish findings based on quadratic equation 5n^2 + 2*x + 1 also either eats up memory or is too slow.
	I've noticed all the solutions these approaches found have the form of F_2k/F_2k+1 for x,
	so it was easy to find solution for k=15, but analysis was missing.

	Analysis:
	https://en.wikipedia.org/wiki/Generating_function
	A_F(x) is the ordinary generating function of the Fibonacci sequence, A_F(x) = x/(1 - x - x^2)
	A_F(x) = sum(1..inf)[x^k*Fk]

	Per user Lakarn and above wiki, following can be derived
	x*A_F(x) = sum(1..inf)[x^(k+1)*Fk] and x^2*A_F(x) = sum(1..inf)[x^(k+2)*Fk], so
		(1 - x - x^2)*A_F(x) = sum(1..inf)[(x^k - x^(k+1) - x^(k+2))*Fk] =
		sum(1..inf)[x^k*Fk] - sum(1..inf)[x^(k+1)*(Fk+1-Fk-1)] - sum(1..inf)[x^(k+2)*(Fk+2-Fk+1)] =
		sum(1..inf)[x^k*Fk] - sum(2..inf)[x^k*(Fk-Fk-2)] - sum(3..inf)[x^k*(Fk-Fk-1)] =
		x*F1 + x^2*F2 + sum(3..inf)[x^k*Fk] - x^2*(F2 - F0) - sum(3..inf)[x^k*(Fk-Fk-2)] - sum(3..inf)[x^k*(Fk-Fk-1)] =
		x + sum(3..inf)[x^k*Fk - x^k*(Fk-Fk-2) - x^k*(Fk-Fk-1)] = x + sum(3..inf)[x^k*Fk-2 - x^k*(Fk-Fk-1)] =
		x + sum(3..inf)[x^k*(Fk-2 - Fk + Fk-1)] = x + sum(3..inf)[x^k]*(Fk-2 - Fk-1 - Fk-2 + Fk-1)] = x

	A_F(x) = x/(1 - x - x^2)	(1)

	According to daniel.is.fischer A_F(-1/x) = A_F(x). Why?

	A_F(-1/x) = -1/x/(1 + 1/x - 1/x^2) = x/(-x^2 - x + 1) = A_F(x)
	So we only need to consider positive x.

	I can't currently prove it, but x = F_2k/F_2k+1 and n = F_2k * F_2k+1

	Alternative approach (with a beautiful correlation to Pell's equation):

	x/(1 - x - x^2) = n => n - nx -nx^2 = x => nx^2 + (n+1)x - n = 0 => x1,2 = (-n-1 +- sqrt((n+1)^2 + 4n^2))/2n
	For x to be rational, 5n^2 + 2n + 1 has to be perfect square
	5n^2 + 2n + 1 = b^2 / *5
	25n^2 + 10n + 5 = 5b^2
	(5n + 1)^2 - 5b^2 = -4

	a^2 - 5b^2 = -4				(2)
	fundamental solution (1, 1, -4)

	(an + bn*sqrt(N))/2 = ((a1 + b1*sqrt(N))/2)^n
	(ak+1 + bk+1*sqrt(N))/2 = ((a1 + b1*sqrt(N))/2)^k * (a1 + b1*sqrt(N))/2 = (ak + bk*sqrt(N))/2 * (a1 + b1*sqrt(N))/2
		= 1/4 * (ak*a1 + bk*a1*sqrt(N) + ak*b1*sqrt(N) + bk*b1*N)

	ak+1 = (ak*a1 + N*bk*b1)/2
	bk+1 = (bk*a1 + ak*b1)/2

	ak+2 = (ak+1*a1 + N*bk+1*b1)/2 = ((ak*a1 + N*bk*b1)/2)*a1 + N*(bk*a1 + ak*b1)/2*b1)/2 =
		(ak*a1^2 + N*bk*b1*a1 + N*(bk*a1*b1 + ak*b1^2))/4 = (ak*(a1^2 + N*b1^2) + 2*N*bk*b1*a1)/4
	bk+2 = (bk+1*a1 + ak+1*b1)/2 = ((bk*a1 + ak*b1)/2*a1 + b1*(ak*a1 + N*bk*b1)/2)/2 =
		(bk*a1^2 + 2*ak*b1*a1 + N*bk*b1^2)/4 = (bk*(a1^2 + N*b1^2) + 2*ak*b1*a1)/4

	ak+2 = (ak*(a1^2 + N*b1^2) + 2*N*bk*b1*a1)/4
	bk+2 = (bk*(a1^2 + N*b1^2) + 2*ak*b1*a1)/4

	ak+2 = (3*ak+ 5*bk)/2
	bk+2 = (3*bk + ak)/2

	1,1			0
	4,2
	11,5		2
	29,13
	76,34		15
	199,89
	521,233		104
*/
