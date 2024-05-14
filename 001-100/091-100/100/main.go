package main

import (
	"math/big"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 100; Arranged Probability
If a box contains twenty-one coloured discs, composed of fifteen blue discs and six red discs, and two discs were taken at random,
it can be seen that the probability of taking two blue discs, P(BB) = (15/21) * (14/20) = 1/2.

The next such arrangement, for which there is exactly 50% chance of taking two blue discs at random, is a box containing
eighty-five blue discs and thirty-five red discs.

By finding the first arrangement to contain over 10^12 = 1,000,000,000,000 discs in total, determine the number of blue discs that the box would contain.
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	// probability of taking two blue discs:
	// b/t * (b - 1)/(t - 1) = 1/2
	// 2*b*(b-1) = t*(t-1)
	// since (t - 1/2)^2 = t^2 - t + 1/4, equation can be written as
	// 2*(b - 1/2)^2 - 1/2 = (t - 1/2)^2 - 1/4 / *4
	// (2*b - 1)^2 - 2 = (2*t - 1)^2 - 1
	// (2*t - 1)^2 - 2*(2*b - 1)^2 = -1, which is a negative Pell's equation
	// 2*t - 1 > 2*10^12 - 1
	// from https://math.stackexchange.com/questions/531833/generating-all-solutions-for-a-negative-pell-equation
	// by user André Nicolas:
	//
	// The fundamental solution of the equation x^2−2*y^2 = −1 is (1,1). We get all positive solutions
	// by taking odd powers of 1 + sqrt(2). The positive solutions are (xn,yn),
	// where xn + yn*sqrt(2)=(1 + sqrt(2))^(2n−1)
	// One can alternately obtain a recurrence for the solutions. If (xn,yn) is a positive solution,
	// then the "next" solution (xn+1,yn+1) is given by
	// xn+1 = 3*xn + 4*yn, yn+1 = 2*xn + 3*yn.
	// Note that your solution (7,5) is the case n=2, and (41,29) is the case n=3.
	// Similarly, the positive solutions of the equation x^2−d*y^2 = −1 (if they exist) are obtained by
	// taking the odd powers of the fundamental solution.

	one := big.NewInt(1)
	two := big.NewInt(2)
	three := big.NewInt(3)
	four := big.NewInt(4)
	x := big.NewInt(1)
	y := big.NewInt(1)
	xp := big.NewInt(1)
	yp := big.NewInt(1)
	temp := big.NewInt(1)
	limit := big.NewInt(2000000000000 - 1)

	for x.Cmp(limit) != 1 {
		x.Mul(xp, three)
		temp.Set(yp)
		temp.Mul(yp, four)
		x.Add(x, temp)

		y.Mul(xp, two)
		temp.Set(yp)
		temp.Mul(yp, three)
		y.Add(y, temp)

		xp.Set(x)
		yp.Set(y)
	}

	y.Add(y, one)
	y.Div(y, two)

	result = strconv.FormatInt(y.Int64(), 10)
	return
}
