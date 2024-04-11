package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 93; Arithmetic Expressions
By using each of the digits from the set, {1, 2, 3, 4}, exactly once, and making use of the four arithmetic operations (+, -, *, )
and brackets/parentheses, it is possible to form different positive integer targets.
For example,
8 = (4 * (1 + 3)) / 2
14 = 4 * (3 + 1 / 2)
19 = 4 * (2 + 3) - 1
36 = 3 * 4 * (2 + 1)
Note that concatenations of the digits, like 12 + 34, are not allowed.
Using the set, {1, 2, 3, 4}, it is possible to obtain thirty-one different target numbers of which 36 is the maximum,
and each of the numbers 1 to 28 can be obtained before encountering the first non-expressible number. Find the set of four distinct digits,
a < b < c < d, for which the longest set of consecutive positive integers, 1 to n, can be obtained, giving your answer as a string: abcd.
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
		limit = 10
	}

	projecteuler.Timed(calc, limit)
}

type (
	opType int
	frac   struct{ up, down int }
)

const (
	plus opType = iota
	minus
	times
	divide
)

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)
	maxConsecutives := 0
	maxAbcd := 0
	opVars := opVariations()

	for a := 1; a < limit-3; a++ {
		for b := a + 1; b < limit-2; b++ {
			for c := b + 1; c < limit-1; c++ {
				for d := c + 1; d < limit; d++ {
					x := consecutives(a, b, c, d, opVars)
					if x > maxConsecutives {
						maxConsecutives = x
						maxAbcd = 1000*a + 100*b + 10*c + d
					}
				}
			}
		}
	}

	result = strconv.Itoa(maxAbcd)

	return
}

func consecutives(a, b, c, d int, opVars [][]opType) int {
	mask := map[byte]int{0: a, 1: b, 2: c, 3: d}
	digitPerms := projecteuler.Permutations(4, nil)
	mNum := map[int]struct{}{}

	for j := range digitPerms {
		currPerm := []int{mask[digitPerms[j][0]], mask[digitPerms[j][1]], mask[digitPerms[j][2]], mask[digitPerms[j][3]]}
		for _, ops := range opVars {
			nums := evalPerm(currPerm, ops)
			for _, x := range nums {
				mNum[x] = struct{}{}
			}
		}
	}

	i := 1
	for ; ; i++ {
		if _, exists := mNum[i]; !exists {
			break
		}
	}

	return i - 1
}

func opVariations() [][]opType {
	ops := make([][]opType, 64)
	for i := 0; i < 64; i++ {
		ops[i] = make([]opType, 3)
	}

	cnt := 0
	for i1 := 0; i1 < 4; i1++ {
		for i2 := 0; i2 < 4; i2++ {
			for i3 := 0; i3 < 4; i3++ {
				ops[cnt][0] = opType(i1)
				ops[cnt][1] = opType(i2)
				ops[cnt][2] = opType(i3)
				cnt++
			}
		}
	}

	return ops
}

func evalPerm(perm []int, ops []opType) []int {
	// ((0 1) 2) 3 -> t0 = 0 1, t1 = t0 2, t2 = t1 3
	// (0 (1 2)) 3 -> t0 = 1 2, t1 = 0 t0, t2 = t1 3
	// 0 ((1 2) 3) -> t0 = 1 2, t1 = t0 3, t2 = 0 t1
	// 0 (1 (2 3)) -> t0 = 2 3, t1 = 1 t0, t2 = 0 t1
	// (0 1) (2 3) -> t0 = 0 1, t1 = 2 3, t2 = t0 t1

	nums := []int{}
	var t0, t1, t2 frac

	// v0 ((0 1) 2) 3 -> t0 = 0 1, t1 = t0 2, t2 = t1 3
	t0 = evalTerm(int2Frac(perm[0]), int2Frac(perm[1]), ops[0])
	t1 = evalTerm(t0, int2Frac(perm[2]), ops[1])
	t2 = evalTerm(t1, int2Frac(perm[3]), ops[2])
	if t2.down != 0 && t2.up*t2.down > 0 && t2.up%t2.down == 0 {
		x := t2.up / t2.down
		nums = append(nums, x)
	}

	// v1 (0 (1 2)) 3 -> t0 = 1 2, t1 = 0 t0, t2 = t1 3
	t0 = evalTerm(int2Frac(perm[1]), int2Frac(perm[2]), ops[1])
	t1 = evalTerm(int2Frac(perm[0]), t0, ops[0])
	t2 = evalTerm(t1, int2Frac(perm[3]), ops[2])
	if t2.down != 0 && t2.up*t2.down > 0 && t2.up%t2.down == 0 {
		x := t2.up / t2.down
		nums = append(nums, x)
	}

	// v2 0 ((1 2) 3) -> t0 = 1 2, t1 = t0 3, t2 = 0 t1
	t0 = evalTerm(int2Frac(perm[1]), int2Frac(perm[2]), ops[1])
	t1 = evalTerm(t0, int2Frac(perm[3]), ops[2])
	t2 = evalTerm(int2Frac(perm[0]), t1, ops[0])
	if t2.down != 0 && t2.up*t2.down > 0 && t2.up%t2.down == 0 {
		x := t2.up / t2.down
		nums = append(nums, x)
	}

	// v3 0 (1 (2 3)) -> t0 = 2 3, t1 = 1 t0, t2 = 0 t1
	t0 = evalTerm(int2Frac(perm[2]), int2Frac(perm[3]), ops[2])
	t1 = evalTerm(int2Frac(perm[1]), t0, ops[1])
	t2 = evalTerm(int2Frac(perm[0]), t1, ops[0])
	if t2.down != 0 && t2.up*t2.down > 0 && t2.up%t2.down == 0 {
		x := t2.up / t2.down
		nums = append(nums, x)
	}

	// v4 (0 1) (2 3) -> t0 = 0 1, t1 = 2 3, t2 = t0 t1
	t0 = evalTerm(int2Frac(perm[0]), int2Frac(perm[1]), ops[0])
	t1 = evalTerm(int2Frac(perm[2]), int2Frac(perm[3]), ops[2])
	t2 = evalTerm(t0, t1, ops[1])
	if t2.down != 0 && t2.up*t2.down > 0 && t2.up%t2.down == 0 {
		x := t2.up / t2.down
		nums = append(nums, x)
	}

	return nums
}

func evalTerm(fa, fb frac, op opType) frac {
	res := frac{}
	switch op {
	case plus:
		res = add(fa, fb)
	case minus:
		res = sub(fa, fb)
	case times:
		res = mul(fa, fb)
	case divide:
		if fb.up != 0 {
			res = div(fa, fb)
		}
	}

	return res
}

func int2Frac(x int) frac {
	return frac{up: x, down: 1}
}

func add(fa, fb frac) frac {
	return frac{up: fa.up*fb.down + fb.up*fa.down, down: fa.down * fb.down}
}

func sub(fa, fb frac) frac {
	return frac{up: fa.up*fb.down - fb.up*fa.down, down: fa.down * fb.down}
}

func mul(fa, fb frac) frac {
	return frac{up: fa.up * fb.up, down: fa.down * fb.down}
}

func div(fa, fb frac) frac {
	return frac{up: fa.up * fb.down, down: fa.down * fb.up}
}
