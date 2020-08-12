package main

import (
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 66; Diophantine equation

Consider quadratic Diophantine equations of the form:
x^2 – Dy^2 = 1

For example, when D=13, the minimal solution in x is 649^2 – 13×180^2 = 1.
It can be assumed that there are no solutions in positive integers when D is square.
By finding minimal solutions in x for D = {2, 3, 5, 6, 7}, we obtain the following:

3^2 – 2×2^2 = 1
2^2 – 3×1^2 = 1
9^2 – 5×4^2 = 1
5^2 – 6×2^2 = 1
8^2 – 7×3^2 = 1

Hence, by considering minimal solutions in x for D <= 7, the largest x is obtained when D = 5.
Find the value of D <= 1000 in minimal solutions of x for which the largest value of x is obtained.
*/

func main() {
	var dLimit int

	if len(os.Args) > 1 {
		dLimit64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		dLimit = int(dLimit64)
	} else {
		dLimit = 1000
	}

	projecteuler.Timed(calc, dLimit, chakravalaSolver{})
}

func calc(args ...interface{}) (result string, err error) {
	dLimit := args[0].(int)
	ms := args[1].(minSolver)

	sqMap := projecteuler.ReverseSquares(dLimit)
	largestMinimalSolution := big.NewInt(0)
	largestD := 0

	for d := 2; d <= dLimit; d++ {
		if _, ok := sqMap[d]; ok {
			continue
		}

		m := ms.minSolve(d)
		if m.Cmp(largestMinimalSolution) == 1 {
			largestMinimalSolution.Set(m)
			largestD = d
		}
	}

	result = strconv.Itoa(largestD)
	return
}

type (
	minSolver interface {
		// minSolve calculates minimal solution to x^2 + d*y^2 = 1
		minSolve(d int) *big.Int
	}

	continuedFractionSolver struct {
		dLimit int
		primes []int
	}

	chakravalaSolver struct{}
)

func makeContinuedFractionSolver(dLimit int) minSolver {
	return continuedFractionSolver{dLimit: dLimit, primes: projecteuler.Primes(dLimit+1, nil)}
}

func (r continuedFractionSolver) minSolve(d int) *big.Int {
	// ROOT CONTINUED FRACTION CONVERGENTS:
	//
	// x^2 – Dy^2 = 1 is Pell's equation.
	// There is a sequence of convergents to the continued fraction for sqrt(D)
	// Each convergent is a rational number, xi/yi where xi and yi are positive integers.
	// Loop through convergents until finding a solution to this particular Pell's equation for (xi,yi).

	dC := projecteuler.MakeContinuedFraction(d, r.primes)
	dBig := big.NewInt(int64(d))

	for i := 1; ; i++ {
		bif := dC.Convergent(i)
		x := big.NewInt(1)
		y := big.NewInt(1)
		x.Set(bif.Numerator())
		y.Set(bif.Denominator())
		x.Mul(x, x)
		y.Mul(y, y)
		y.Mul(y, dBig)
		x.Sub(x, y)
		if x.String() == "1" {
			return bif.Numerator()
		}
	}
}

func (c chakravalaSolver) minSolve(d int) *big.Int {
	return projecteuler.Chakravala(d).X
}
