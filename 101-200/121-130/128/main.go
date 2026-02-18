package main

import (
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 128; Hexagonal Tile Differences
A hexagonal tile with number 1 is surrounded by a ring of six hexagonal tiles, starting at "12 o'clock"
and numbering the tiles 2 to 7 in an anti-clockwise direction.

New rings are added in the same fashion, with the next rings being numbered 8 to 19, 20 to 37, 38 to 61, and so on.
The diagram below shows the first three rings.

By finding the difference between tile n and each of its six neighbours we shall define PD(n)
to be the number of those differences which are prime.

For example, working clockwise around tile 8 the differences are 12, 29, 11, 6, 1, and 13. So PD(8) = 3.
In the same way, the differences around tile 17 are 1, 17, 16, 1, 11, and 10, hence PD(17) = 2.
It can be shown that the maximum value of PD(n) is 3.

If all of the tiles for which PD(n) = 3 are listed in ascending order to form a sequence, the 10th tile would be 271.
Find the 2000th tile in this sequence.
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
		limit = 2000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primes := projecteuler.Primes(1000000, nil)

	res := make([]int, 0, limit)
	res = append(res, 1)
	res = append(res, 2)
	res = append(res, 8)

	count := 3
	prevCube := 8
	x := 0
	for i := 3; ; i++ {
		cube := i * i * i

		// ultimate
		x = cube - prevCube
		diffsToExamine := []int{12*i - 19, 6*i - 7, 6*i - 1}
		currCount := 0
		for j := 0; j < 3; j++ {
			if _, found := slices.BinarySearch(primes, diffsToExamine[j]); found {
				currCount++
			}
		}
		if currCount == 3 {
			count++
			res = append(res, x)
		}
		if count == limit {
			break
		}

		// top
		x++
		diffsToExamine = []int{6*i - 1, 6*i + 1, 12*i + 5}
		currCount = 0
		for j := 0; j < 3; j++ {
			if _, found := slices.BinarySearch(primes, diffsToExamine[j]); found {
				currCount++
			}
		}
		if currCount == 3 {
			count++
			res = append(res, x)
		}
		if count == limit {
			break
		}

		prevCube = cube
	}

	result = strconv.Itoa(res[limit-1])
	return
}

/*
	// top
	x - t.cardinality[ring-1]
	x + 1
	t.max[ring]
	t.max[ring] + 1
	t.max[ring] + 2
	t.max[ring+1]

	DIFFS
	t.cardinality[ring-1]=6*(i-1)				no
	1											no
	-i^3+(i-1)^3-1+(i+1)^3-i^3=
		(i+1)^3-2i^3+(i-1)^3-1=
		i^3+3i^2+3i+1-2i^3+i^3-3i^2+3i-1-1=
		3i^2+3i+1-3i^2+3i-2=6i-1				yes
		6i										no
		6i+1									yes
		(i+2)^3-(i+1)^3-i^3+(i-1)^3-1=
		i^3 + 6i^2 + 12i + 8 - (i^3 + 3i^2 + 3i + 1 + i^3 - i^3 + 3i^2 - 3i + 1 + 1)=
		6i^2 + 12i + 8 - (3i^2 + 3i + 1 + 3i^2 - 3i + 1 + 1)=
		12i + 5									yes

	// ultimate
	if x == t.max[ring] {
	t.corners[ring-1][0]
	t.max[ring-1]
	t.corners[ring][0]
	x - 1
	t.max[ring+1] - 1
	t.max[ring+1]

	x == cubes[i+1]-cubes[i] = 64-27=37
	cubes[i-1]-cubes[i-2]+1
	cubes[i]-cubes[i-1]
	cubes[i]-cubes[i-1]+1
	x - 1
	cubes[i+2]-cubes[i+1] - 1
	cubes[i+2]-cubes[i+1]

	DIFFS
	cubes[i+1]-cubes[i]-cubes[i-1]+cubes[i-2]-1=
		i^3 + 3i^2 + 3i + 1 - i^3 - (i^3 - 3i^2 + 3i - 1) + (i^3 - 6i^2 + 12i - 8) - 1=
		3i^2 + 3i + 1 + i^3 - 6i^2 + 12i - 8 - i^3 + 3i^2 - 3i + 1 - 1=
		3i^2 + 3i + 1 - 6i^2 + 12i + 3i^2 - 3i - 8=
		12i - 7									yes
	cubes[i+1]-cubes[i]-cubes[i]+cubes[i-1]=
		6i										no
	cubes[i+1]-cubes[i]-cubes[i]+cubes[i-1]-1=
		6i - 1									yes
	1											no
	cubes[i+1]-cubes[i]-cubes[i+2]+cubes[i+1]+1=
		cubes[i+2]-2cubes[i+1]+cubes[i]-1=
		i^3 + 6i^2 + 12i + 8 -2(i^3 + 3i^2 + 3i + 1)+i^3-1=
		12i+8-6i-2-1=
		6i+5									yes
		6i+6									no

	// corners
	if (x-t.max[ring-1]-1)%ring == 0 {
		ret[0] = x - t.cardinality[ring-1] - side
		ret[1] = x - 1
		ret[2] = x + 1
		ret[3] = t.corners[ring+1][side] - 1
		ret[4] = t.corners[ring+1][side]
		ret[5] = t.corners[ring+1][side] + 1

	DIFFS
	x == t.max[ring-1] + j*ring + 1	(j=1..5)
	t.cardinality[ring-1] + side
	1
	1
	t.corners[ring+1][side] - t.max[ring] - j*ring - 1
	t.corners[ring+1][side] - t.max[ring] - j*ring
	t.corners[ring+1][side] - t.max[ring] - j*ring + 1

	6*(i-1) + j [6i-5, 6i-4, 6i-3, 6i-2, 6i-1]	yes for corners 1 and 5
	1											no
	1											no
	(i+1)^3 - i^3 + j*(i+1) + 1 -(i^3 - (i-1)^3 + j*i + 1)=
	3i^2+3i+1 + j*(i+1) + 1 - i^3 + i^3 - 3i^2 + 3i - 1 - j*i - 1=
	3i^2+3i+1 + j + 1 - 3i^2 + 3i-2=
	6i+j-1 [6i,6i+1,6i+2,6i+3,6i+4]				yes for corner 2
	6i+j   [6i+1,6i+2,6i+3,6i+4,6i+5]			yes for corners 1 and 5
	6i+j+1 [6i+2,6i+3,6i+4,6i+5,6i+6]			yes for corner 4
	no need to examine corners

	// regular
	ret[0] = x - t.cardinality[ring-1] - side - 1
	ret[1] = x - t.cardinality[ring-1] - side
	ret[2] = x - 1
	ret[3] = x + 1
	ret[4] = x + t.cardinality[ring] + side
	ret[5] = x + t.cardinality[ring] + side + 1

	DIFFS
	t.cardinality[ring-1] + side + 1
	t.cardinality[ring-1] + side
	1
	1
	t.cardinality[ring] + side
	t.cardinality[ring] + side + 1

	6*(i-1) + j + 1	mod 6 = 2,3,4,5,0			yes for side 4
	6*(i-1) + j 	mod 6 = 1,2,3,4,5			yes for sides 1 and 5
	1											no
	1											no
	6*i + j			mod 6 = 1,2,3,4,5			yes for sides 1 and 5
	6*i + j + 1		mod 6 = 2,3,4,5,0			yes for side 4
	no need to examine regulars

	EXAMINE
	top:				6i-1, 6i+1, 12i+5	for x == i^3 - (i-1)^3 + 1
	ultimate:			12i-7, 6i-1, 6i+5	for x == (i+1)^3 - i^3
*/
