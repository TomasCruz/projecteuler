package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 72; Counting Fractions
Consider the fraction, n/d, where n and d are positive integers. If n<d and HCF(n,d)=1, it is called a reduced proper fraction.
If we list the set of reduced proper fractions for d<=8 in ascending order of size, we get:
1/8,1/7,1/6,1/5,1/4,2/7,1/3,3/8,2/5,3/7,1/2,4/7,3/5,5/8,2/3,5/7,3/4,4/5,5/6,6/7,7/8
It can be seen that there are 21 elements in this set.
How many elements would be contained in the set of reduced proper fractions for d <= 1000000?
?*/

func main() {
	var limit int

	if len(os.Args) > 1 {
		limit64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limit = int(limit64)
	} else {
		limit = 1000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	// count := 0
	// m := buildFractionMatrix(limit)
	// for d := 0; d <= limit; d++ {
	// 	for n := 0; n < limit; n++ {
	// 		if m[d][n] == 1 {
	// 			count++
	// 		}
	// 	}
	// }

	count := int64(0)
	primes := projecteuler.Primes(limit, nil)
	for d := 2; d <= limit; d++ {
		// lesser := relativelyPrimesLessThan(d, primes)
		// for _, x := range lesser {
		// 	fmt.Printf("%d/%d, ", x, d)
		// }
		// fmt.Println()

		// 30397485
		// 303963552391
		// 303963552391

		// count += int64(len(lesser))
		count += int64(projecteuler.Totient(d, primes)) //relativelyPrimesLessThan(d, primes))
	}

	// fractionSet := map[projecteuler.Fraction]struct{}{}
	// for d := 2; d <= limit; d++ {
	// 	for n := 1; n < d; n++ {
	// 		f := projecteuler.NewFraction(int64(n), int64(d))
	// 		f.Reduce()
	// 		fractionSet[f] = struct{}{}
	// 	}
	// }

	// var fs []projecteuler.Fraction
	// for d := 2; d <= limit; d++ {
	// 	for n := 1; n < d; n++ {
	// 		if m[d][n] == 1 {
	// 			f := projecteuler.NewFraction(n, d)
	// 			f.Reduce()
	// 			fs = append(fs, f)
	// 		}
	// 	}
	// }

	// threeSevenths := projecteuler.NewFraction(3, 7)
	// i := 1
	// for ; i < len(fs); i++ {
	// 	if fs[i] == threeSevenths {
	// 		break
	// 	}
	// }

	// fs := make([]projecteuler.Fraction, limit+1)
	// for d, m := 3, 3; d <= limit; d++ {
	// 	// x/d < 3/7 => x < 3*d/7
	// 	x := (3 * d) / 7
	// 	if m == 7 {
	// 		m = 0
	// 		x -= 1
	// 	}

	// 	fs[d] = projecteuler.NewFraction(int64(x), int64(d))
	// 	m++
	// }
	// // sort.Sort(projecteuler.FractionSlice(fs))

	// max := projecteuler.NewFraction(int64(1), int64(3))
	// for d := 3; d <= limit; d++ {
	// 	if max.Less(fs[d]) {
	// 		max = fs[d]
	// 	}
	// }

	result = strconv.FormatInt(count, 10)
	// result = strconv.Itoa(count)
	return
}

func buildFractionMatrix(limit int) [][]int {
	m := make([][]int, limit+1) // d = 3..limit
	m[0] = make([]int, limit)
	m[1] = make([]int, limit)
	for d := 2; d <= limit; d++ {
		m[d] = make([]int, limit) // n = 1..limit-1
		for n := 1; n < d; n++ {
			m[d][n] = 1
		}
	}

	for d := 2; d <= limit; d++ {
		for n := 1; n < d; n++ {
			for k := 2; n*k < limit && d*k <= limit; k++ {
				m[d*k][n*k] = 0
			}
		}
	}

	// for d := 0; d <= limit; d++ {
	// 	for n := 0; n < limit; n++ {
	// 		fmt.Printf("%d ", m[d][n])
	// 	}
	// 	fmt.Println()
	// }

	return m
}

func relativelyPrimesLessThan(d int, primes []int) int {
	var divisors []int
	i := 0
	l := len(primes)
	for i < l && primes[i] < d {
		if d%primes[i] == 0 {
			divisors = append(divisors, primes[i])
		}
		i++
	}

	multiples := map[int]struct{}{}
	for i := 0; i < len(divisors); i++ {
		k := 1
		for k*divisors[i] < d {
			multiples[k*divisors[i]] = struct{}{}
			k++
		}
	}

	// var result []int
	// for i := 1; i < d; i++ {
	// 	if _, present := multiples[i]; !present {
	// 		result = append(result, i)
	// 	}
	// }

	// return result
	return d - 1 - len(multiples)
}
