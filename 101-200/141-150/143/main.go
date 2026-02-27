package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 143; Torricelli Triangles
Let ABC be a triangle with all interior angles being less than 120 degrees. Let X be any point inside the triangle and let
XA = p, XC = q, and XB = r.

Fermat challenged Torricelli to find the position of X such that p + q + r was minimised.

Torricelli was able to prove that if equilateral triangles AOB, BNC and AMC are constructed on each side of triangle ABC,
the circumscribed circles of AOB, BNC, and AMC will intersect at a single point, T, inside the triangle.
Moreover he proved that T, called the Torricelli/Fermat point, minimises p + q + r. Even more remarkable, it can be shown that
when the sum is minimised, AN = BM = CO = p + q + r and that AN, BM and CO also intersect at T.

If the sum is minimised and a, b, c, p, q and r are all positive integers we shall call triangle ABC a Torricelli triangle.
For example, a = 399, b = 455, c = 511 is an example of a Torricelli triangle, with p + q + r = 784.

Find the sum of all distinct values of p + q + r <= 120000 for Torricelli triangles.
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
		limit = 120000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	shortSet := map[int]map[[3]int]struct{}{}
	for m, modM := 2, 2; m < limit; m, modM = m+1, modM+1 {
		if modM == 3 {
			modM = 0
		}

		for n, modN := 1, 1; n < m; n, modN = n+1, modN+1 {
			if modN == 3 {
				modN = 0
			}

			if modN == modM {
				continue
			}

			s1 := 2*m*n + n*n
			s2 := m*m - n*n
			s3 := m*m + m*n + n*n

			sum := s1 + s2
			if sum >= limit {
				break
			}

			t := [3]int{s1, s2, s3}
			if s1 > s2 {
				t = [3]int{s2, s1, s3}
			}

			if _, present := shortSet[t[0]]; !present {
				shortSet[t[0]] = map[[3]int]struct{}{}
			}
			shortSet[t[0]][t] = struct{}{}

			addMultiples(t, limit, shortSet)
		}
	}

	pqrSet := combinePQRs(limit, shortSet)

	sum := int64(0)
	for pqr := range pqrSet {
		sum += int64(pqr)
	}

	result = strconv.FormatInt(sum, 10)
	return
}

func addMultiples(t [3]int, limit int, shortSet map[int]map[[3]int]struct{}) {
	for k := 1; ; k++ {
		v := [3]int{k * t[0], k * t[1], k * t[2]}
		if !vecOK(limit, v) {
			break
		}

		if _, present := shortSet[v[0]]; !present {
			shortSet[v[0]] = map[[3]int]struct{}{}
		}
		shortSet[v[0]][v] = struct{}{}
	}
}

func vecOK(limit int, v [3]int) bool {
	p := v[0] + v[1]

	if p >= limit {
		return false
	}

	return true
}

func combinePQRs(limit int, shortSet map[int]map[[3]int]struct{}) map[int]map[[7]int]struct{} {
	ret := map[int]map[[7]int]struct{}{}

	for s1 := range shortSet {
		for t1 := range shortSet[s1] {
			for t2 := range shortSet[s1] {
				if t2[1] < t1[1] {
					continue
				}

				shorter, longer := t1[1], t2[1]
				for t3 := range shortSet[shorter] {
					if t3[1] != longer {
						continue
					}

					a := t1[2]
					b := t2[2]
					c := t3[2]
					if a+b <= c || a+c <= b || b+c <= a {
						continue
					}

					pqr := s1 + shorter + longer
					if pqr >= limit {
						continue
					}

					if _, present := ret[pqr]; !present {
						ret[pqr] = map[[7]int]struct{}{}
					}

					arr := [7]int{s1, shorter, longer, pqr, a, b, c}
					ret[pqr][arr] = struct{}{}
				}
			}
		}
	}

	return ret
}
