package main

import (
	"math"
)

type minPathMatrix struct {
	limit int
	input [][]int64
	m     [][]int64
}

func newMinPathMatrix(limit int, input [][]int64) minPathMatrix {
	return minPathMatrix{
		limit: limit,
		input: input,
		m:     makeMatrix(limit),
	}
}

func (mm minPathMatrix) buildMinPathMatrix() {
	for i := 0; i < mm.limit; i++ {
		for j := 0; j < mm.limit; j++ {
			mm.m[i][j] = math.MaxInt64
		}
	}
	mm.m[0][0] = mm.input[0][0]

	for i := 0; i < mm.limit; i++ {
		for j := 0; j < mm.limit; j++ {
			fromLeft := int64(math.MaxInt64)
			fromUp := int64(math.MaxInt64)

			if mm.canGo(i, j, left) {
				leftR, leftC := mm.goDirection(i, j, left)
				fromLeft = mm.m[leftR][leftC] + mm.input[i][j]
			}
			if mm.canGo(i, j, up) {
				upR, upC := mm.goDirection(i, j, up)
				fromUp = mm.m[upR][upC] + mm.input[i][j]
			}

			minFrom := fromUp
			if fromLeft < fromUp {
				minFrom = fromLeft
			}

			if minFrom < mm.m[i][j] {
				mm.m[i][j] = minFrom
			}
		}
	}
}

func (mm minPathMatrix) canGo(i, j int, dir direction) bool {
	switch dir {
	case left:
		return j > 0
	case down:
		return i < mm.limit-1
	case right:
		return j < mm.limit-1
	case up:
		return i > 0
	}

	return false
}

func (mm minPathMatrix) goDirection(r, c int, dir direction) (nr, nc int) {
	nr = r
	nc = c

	switch dir {
	case left:
		nc--
	case down:
		nr++
	case right:
		nc++
	case up:
		nr--
	}

	return
}
