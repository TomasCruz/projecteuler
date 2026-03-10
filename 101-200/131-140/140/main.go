package main

import (
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 140; Modified Fibonacci Golden Nuggets
Consider the infinite polynomial series A_G(x) = x*G1 + x^2*G2 + x^3*G3 + ...,
where Gk is the kth term of the second order recurrence relation Gk = Gk-1 + Gk-2,
G_1 = 1 and G_2 = 4; that is, 1, 4, 5, 9, 14, 23, ...

For this problem we shall be interested in values of x for which A_G(x) is a positive integer.

The corresponding values of x for the first five natural numbers are shown below.

x	A_G(x)</th>
-------------------
(sqrt(5)-1)/4		1
2/5					2
(sqrt(22)-2)/6		3
(sqrt(137)-5)/14	4
1/2					5

We shall call A_G(x) a golden nugget if x is rational, because they become increasingly rarer; for example, the 20th golden nugget is 211345365.

Find the sum of the first thirty golden nuggets.
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
		limit = 30
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	one, fundamentals := projecteuler.LMM(5, 44)

	nugSet := map[projecteuler.PellTriplet]struct{}{}
	solutions := map[projecteuler.PellTriplet]struct{}{}
	for _, f := range fundamentals {
		solutions[f] = struct{}{}
		if f.A%5 == 2 && f.A > 7 {
			nugSet[f] = struct{}{}
		}
	}

	for len(nugSet) < int(1.5*float64(limit)) {
		for s := range solutions {
			t := projecteuler.ComposePellTriplets(5, one, s)
			solutions[t] = struct{}{}
			if t.A%5 == 2 && t.A > 7 {
				nugSet[t] = struct{}{}
			}
		}
	}

	nugSlice := make([]projecteuler.PellTriplet, 0, len(nugSet))
	for n := range nugSet {
		nugSlice = append(nugSlice, n)
	}
	slices.SortFunc(nugSlice, projecteuler.TripletSortFunc)
	nugSlice = nugSlice[:limit]

	sum := int64(0)
	nuggets := make([]int64, limit)
	for i := range limit {
		num := (nugSlice[i].A - 7) / 5
		nuggets[i] = num
		sum += num
	}

	result = strconv.FormatInt(sum, 10)
	return
}

/*
	x*A_G(x) = sum(1..inf)[x^(k+1)*Gk] and x^2*A_G(x) = sum(1..inf)[x^(k+2)*Gk], so
	(1 - x - x^2)*A_G(x) = sum(1..inf)[(x^k - x^(k+1) - x^(k+2))*Gk] =
		sum(1..inf)[x^k*Gk] - sum(1..inf)[x^(k+1)*(Gk+2 - Gk+1)] - sum(1..inf)[x^(k+2)*(Gk+2 - Gk+1)] =
		x*G1 + x^2*G2 + sum(3..inf)[x^k*Gk] - x^2*(G3-G2) - sum(3..inf)[x^k*(Gk+1-Gk)] - sum(3..inf)[x^k*(Gk-Gk-1)] =
		x + 3*x^2 + sum(3..inf)[x^k*(Gk - Gk+1 + Gk - Gk + Gk-1)] = x + 3*x^2

	A_G(x) = (x + 3*x^2)/(1 - x - x^2)	(1)

	(x + 3*x^2)/(1 - x - x^2) = n => n - nx - nx^2 = x + 3*x^2 => (n + 3)x^2 + (n+1)x - n = 0 => x1,2 = (-n-1 +- sqrt((n+1)^2 + 4n*(n+3)))/2*(n+3)
	For x to be rational, 5n^2 + 14n + 1 has to be perfect square

	5n^2 + 14n + 1 = b^2				(2)

	5n^2 + 14n + 1 = b^2 / *5
	25n^2 + 70n + 5 = 5b^2
	(5n + 7)^2 - 5b^2 = 44				(3)
*/
