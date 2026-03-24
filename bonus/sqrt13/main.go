package main

import (
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem sqrt(13); Square root convergents
The decimal expansion of the square root of two is 1.{4142135623}730...

If we define S(n, d) to be the the sum of the first d digits in the fractional part of the decimal expansion of sqrt{n},
it can be seen that S(2, 10) = 4 + 1 + 4 + ... + 3 = 31.

It can be confirmed that S(2, 100) = 481.

Find S(13, 1000).

Note: Instead of just using arbitrary precision floats, try to be creative with your method.
*/

func main() {
	var a1, a2 int

	if len(os.Args) > 1 {
		a1_64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		a1 = int(a1_64)

		a2_64, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		a2 = int(a2_64)
	} else {
		a1 = 13
		a2 = 1000
	}

	projecteuler.Timed(calc, a1, a2)
}

func calc(args ...interface{}) (result string, err error) {
	num := args[0].(int)
	decimalCount := args[1].(int)

	decimals := sqrtDecimalExpansion(num, decimalCount*2)

	sum := 0
	for i := 1; i <= decimalCount; i++ {
		sum += int(decimals[i])
	}

	result = strconv.Itoa(sum)
	return
}

func sqrtDecimalExpansion(num, digitCount int) []byte {
	ret := make([]byte, 0, digitCount)

	period64, matrix := projecteuler.PQa(0, 1, int64(num))
	period := int(period64)
	first := matrix[0][0]
	denominatorSequence := matrix[0][1:]

	p := make([]*big.Int, 0, digitCount)
	p = append(p, big.NewInt(first))
	p = append(p, big.NewInt(matrix[1][1]))
	q := make([]*big.Int, 0, digitCount)
	q = append(q, big.NewInt(1))
	q = append(q, big.NewInt(matrix[2][1]))

	denomCount := 0
	for i := 2; i < digitCount; i++ {
		denomCount++
		if denomCount == period {
			denomCount = 0
		}

		// pk = ak*pk-1 + pk-2
		// qk = ak*qk-1 + qk-2
		ak := big.NewInt(denominatorSequence[denomCount])
		pm := new(big.Int).Mul(ak, p[i-1])
		qm := new(big.Int).Mul(ak, q[i-1])
		p = append(p, pm.Add(pm, p[i-2]))
		q = append(q, qm.Add(qm, q[i-2]))
	}

	rat := new(big.Rat).SetFrac(p[len(p)-1], q[len(p)-1])
	ratString := rat.FloatString(digitCount/2 + 3)

	ret = append(ret, byte(first))
	for i := 2; i < digitCount/2+3; i++ {
		ret = append(ret, []byte(ratString)[i]-'0')
	}

	return ret
}
