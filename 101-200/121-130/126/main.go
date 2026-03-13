package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 126; Cuboid Layers
The minimum number of cubes to cover every visible face on a cuboid measuring 3 * 2 * 1 is twenty-two.

If we then add a second layer to this solid it would require forty-six cubes to cover every visible face,
the third layer would require seventy-eight cubes, and the fourth layer would require one-hundred and eighteen cubes to cover every visible face.

However, the first layer on a cuboid measuring 5 * 1 * 1 also requires twenty-two cubes;
similarly the first layer on cuboids measuring 5 * 3 * 1, 7 * 2 * 1, and 11 * 1 * 1 all contain forty-six cubes.

We shall define C(n) to represent the number of cuboids that contain n cubes in one of its layers. So C(22) = 2, C(46) = 4, C(78) = 5, and C(118) = 8.
It turns out that 154 is the least value of n for which C(n) = 10.

Find the least value of n for which C(n) = 1000.
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
		limit = 1000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	maxCubes := 20 * limit
	cubeCount := make([]int, maxCubes)

	for a := 1; firstLayer(a, 1, 1) < 20*limit; a++ {
		for b := 1; b <= a && firstLayer(a, b, 1) < 20*limit; b++ {
			for c := 1; c <= b && firstLayer(a, b, c) < 20*limit; c++ {
				for i := 1; ; i++ {
					lay := 2*(a*b+a*c+b*c) + 4*(i-1)*(a+b+c+i-2)
					if lay >= maxCubes {
						break
					}

					cubeCount[lay]++
				}
			}
		}
	}

	i := 1
	m := -1
	for ; i < len(cubeCount); i++ {
		if cubeCount[i] > m {
			m = cubeCount[i]
		}

		if cubeCount[i] == limit {
			m = i
			break
		}
	}

	result = strconv.Itoa(m)
	return
}

func firstLayer(a, b, c int) int {
	return 2 * (a*b + a*c + b*c)
}
