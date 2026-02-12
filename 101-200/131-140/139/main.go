package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 139; Pythagorean Tiles
Let (a, b, c) represent the three sides of a right angle triangle with integral length sides.
It is possible to place four such triangles together to form a square with length c.

For example, (3, 4, 5) triangles can be placed together to form a 5 by 5 square with a 1 by 1 hole in the middle
and it can be seen that the 5 by 5 square can be tiled with twenty-five 1 by 1 squares.

However, if (5, 12, 13) triangles were used then the hole would measure 7 by 7 and these could not be used to tile the 13 by 13 square.

Given that the perimeter of the right triangle is less than one-hundred million,
how many Pythagorean triangles would allow such a tiling to take place?
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
		limit = 8
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	power := int64(1)
	for i := 0; i < limit; i++ {
		power *= 10
	}

	// https://en.wikipedia.org/wiki/Tree_of_primitive_Pythagorean_triples
	t := buildTree(power)
	triplets := [][3]int64{}
	triplets = treeSearch(t, triplets, func(v [3]int64) bool {
		bMa := v[1] - v[0]

		if bMa < 0 {
			bMa = -bMa
		}

		return v[2]%bMa == 0
	})

	n := int64(0)
	for _, t := range triplets {
		n += countTriplets(t, power)
	}

	result = strconv.FormatInt(n, 10)
	return
}

type tern struct {
	v        [3]int64
	children [3]*tern
}

func buildTree(power int64) *tern {
	t := &tern{
		v: [3]int64{3, 4, 5},
	}

	nextGen(power, t)
	return t
}

var a = [3][3]int64{
	{1, -2, 2},
	{2, -1, 2},
	{2, -2, 3},
}

var b = [3][3]int64{
	{1, 2, 2},
	{2, 1, 2},
	{2, 2, 3},
}

var c = [3][3]int64{
	{-1, 2, 2},
	{-2, 1, 2},
	{-2, 2, 3},
}

func nextGen(power int64, t *tern) {

	childA := mulMat(a, t.v)
	childB := mulMat(b, t.v)
	childC := mulMat(c, t.v)

	if vecOK(power, childA) {
		t.children[0] = &tern{v: childA}
		nextGen(power, t.children[0])
	}

	if vecOK(power, childB) {
		t.children[1] = &tern{v: childB}
		nextGen(power, t.children[1])
	}

	if vecOK(power, childC) {
		t.children[2] = &tern{v: childC}
		nextGen(power, t.children[2])
	}
}

func mulMat(m [3][3]int64, v [3]int64) [3]int64 {
	res := [3]int64{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res[i] += m[i][j] * v[j]
		}
	}

	return res
}

func vecOK(power int64, v [3]int64) bool {
	p := v[0] + v[1] + v[2]

	if p >= power {
		return false
	}

	return true
}

func treeSearch(t *tern, ret [][3]int64, f func(v [3]int64) bool) [][3]int64 {
	if f(t.v) {
		ret = append(ret, t.v)
	}

	children := []*tern{}
	for i := 0; i < 3; i++ {
		if t.children[i] != nil {
			children = append(children, t.children[i])
		}
	}

	for _, c := range children {
		ret = treeSearch(c, ret, f)
	}

	return ret
}

func countTriplets(t [3]int64, power int64) int64 {
	p := t[0] + t[1] + t[2]

	n := int64(0)
	for i := int64(1); ; i++ {
		if i*p >= power {
			break
		}
		n++
	}

	return n
}
