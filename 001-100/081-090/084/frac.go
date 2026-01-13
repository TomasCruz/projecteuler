package main

import (
	"fmt"
)

type frac struct {
	num, denom int64
}

func (f frac) add(s frac) frac {
	var up, down int64

	if s.denom == f.denom {
		up = f.num + s.num
		down = f.denom
	} else {
		up = f.num*s.denom + s.num*f.denom
		down = f.denom * s.denom
	}

	g := gcd(up, down)
	return frac{num: up / g, denom: down / g}
}

func (f frac) mul(s frac) frac {
	up := f.num * s.num
	down := f.denom * s.denom

	g := gcd(up, down)
	return frac{num: up / g, denom: down / g}
}

func (f frac) toString() string {
	return fmt.Sprintf("%d/%d", f.num, f.denom)
}

func gcd(x, y int64) int64 {
	divisor := int64(2)
	greatestCommonDivisor := int64(1)
	for {
		if divisor > x || divisor > y || divisor > 3 {
			break
		}

		for {
			if x%divisor != 0 || y%divisor != 0 {
				break
			}

			x, y = x/divisor, y/divisor
			greatestCommonDivisor *= divisor
		}

		if x == 1 || y == 1 {
			break
		}
		divisor++
	}

	return greatestCommonDivisor
}

// 2,12 1
// 3,11 2
// 4,10 3
// 5,9 4
// 6,8 5
// 7 6
//
// 2,8 1
// 3,7 2
// 4,6 3
// 5 4
func buildProbabilitiesSingle(limit int) []frac {
	limitSquare := limit * limit
	ret := make([]frac, 2*limit+1)

	ret[0] = frac{num: 0, denom: 1}
	ret[1] = frac{num: 0, denom: 1}
	ret[2] = frac{num: 0, denom: 1}
	ret[2*limit] = frac{num: 0, denom: 1}
	for i := 3; i <= limit; i += 2 {
		ret[i] = frac{num: int64(i - 1), denom: int64(limitSquare)}
		ret[2*limit-i+2] = frac{num: int64(i - 1), denom: int64(limitSquare)}
	}
	for i := 4; i <= limit; i += 2 {
		ret[i] = frac{num: int64(i - 2), denom: int64(limitSquare)}
		ret[2*limit-i+2] = frac{num: int64(i - 2), denom: int64(limitSquare)}
	}
	ret[limit+1] = frac{num: int64(limit), denom: int64(limitSquare)}

	// for i := 0; i <= 2*limit; i++ {
	// 	fmt.Println(ret[i].toString())
	// }
	return ret
}
