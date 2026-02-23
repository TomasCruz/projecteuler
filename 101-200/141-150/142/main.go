package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 142; Perfect Square Collection
Find the smallest x + y + z with integers x > y > z > 0 such that x + y, x - y, x + z, x - z, y + z, y - z are all perfect squares.
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
		limit = 4
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
	legSet, hypoSet := buildSets(t, power)
	sq := findSquares(legSet, hypoSet)

	result = strconv.FormatInt(sq[9], 10)
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

func buildSets(t *tern, power int64) (map[int64]map[[3]int64]struct{}, map[int64]map[[3]int64]struct{}) {
	legSet := map[int64]map[[3]int64]struct{}{}
	hypoSet := map[int64]map[[3]int64]struct{}{}

	buildSetsRec(t, power, legSet, hypoSet)
	return legSet, hypoSet
}

func buildSetsRec(t *tern, power int64, legSet, hypoSet map[int64]map[[3]int64]struct{}) {
	addMultiples(t.v, power, legSet, hypoSet)

	for _, ch := range t.children {
		if ch != nil {
			buildSetsRec(ch, power, legSet, hypoSet)
		}
	}
}

func addMultiples(vec [3]int64, power int64, legSet, hypoSet map[int64]map[[3]int64]struct{}) {
	vOrdered := [3]int64{}
	copy(vOrdered[:], vec[:])
	if vec[0] > vec[1] {
		vOrdered = [3]int64{vec[1], vec[0], vec[2]}
	}

	for k := int64(1); ; k++ {
		v := [3]int64{k * vOrdered[0], k * vOrdered[1], k * vOrdered[2]}
		if !vecOK(power, v) {
			break
		}

		if _, present := legSet[v[0]]; !present {
			legSet[v[0]] = map[[3]int64]struct{}{v: {}}
		} else {
			legSet[v[0]][v] = struct{}{}
		}

		if _, present := legSet[v[1]]; !present {
			legSet[v[1]] = map[[3]int64]struct{}{v: {}}
		} else {
			legSet[v[1]][v] = struct{}{}
		}

		if _, present := hypoSet[v[2]]; !present {
			hypoSet[v[2]] = map[[3]int64]struct{}{v: {}}
		} else {
			hypoSet[v[2]][v] = struct{}{}
		}
	}
}

func findSquares(legSet, hypoSet map[int64]map[[3]int64]struct{}) [10]int64 {
	solutions := [][10]int64{}

	for a, aecSet := range legSet {
		if len(aecSet) < 2 {
			continue
		}

		for aec := range aecSet {
			e := aec[1]
			if e == a {
				e = aec[0]
			}
			c := aec[2]

			for vAec := range aecSet {
				if (vAec[0] == a && vAec[1] == e) || (vAec[0] == e && vAec[1] == a) {
					continue
				}

				f := vAec[0]
				if f == a {
					f = vAec[1]
				}
				d := vAec[2]

				for deb := range legSet[e] {
					if deb[1] != d {
						continue
					}

					b := deb[2]
					_, present1 := hypoSet[b][[3]int64{f, c, b}]
					_, present2 := hypoSet[b][[3]int64{c, f, b}]
					if !present1 && !present2 {
						continue
					}

					x := a*a + b*b
					if x%2 == 1 {
						continue
					}
					x /= 2

					y := (b*b - a*a) / 2

					z := d*d - c*c
					if z%2 == 1 {
						continue
					}
					z /= 2

					solutions = append(solutions, [10]int64{a, b, c, d, e, f, x, y, z, x + y + z})
				}
			}
		}
	}

	m := int64(math.MaxInt64)
	v := [10]int64{}
	for _, s := range solutions {
		if s[9] < m {
			m = s[9]
			v = s
		}
	}

	return v
}
