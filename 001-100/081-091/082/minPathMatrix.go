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

func (mm minPathMatrix) buildMinPathMatrix(r, c int) {
	for i := 0; i < mm.limit; i++ {
		for j := 0; j < mm.limit; j++ {
			mm.m[i][j] = math.MaxInt64
		}
	}
	mm.m[r][c] = mm.input[r][c]

	for i := r; i > 0; i-- {
		mm.m[i-1][0] = mm.m[i][0] + mm.input[i-1][0]
	}

	mm.buildDownRightMatrix()

	for i := 0; i < mm.limit; i++ {
		for j := 0; j < mm.limit; j++ {
			if mm.canGo(i, j, up) {
				chechRes, newMin := mm.checkNewMin(i, j, up)
				if chechRes {
					mm.updateCell(i, j, up, newMin)
				}
			}
		}
	}
}

func (mm minPathMatrix) buildDownRightMatrix() {
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

func (mm minPathMatrix) checkNewMin(r, c int, goingTo direction) (bool, int64) {
	if !mm.canGo(r, c, goingTo) {
		return false, 0
	}

	nr, nc := mm.goDirection(r, c, goingTo)
	newMin := mm.m[r][c] + mm.input[nr][nc]

	return newMin < mm.m[nr][nc], newMin
}

func (mm minPathMatrix) updateCell(r, c int, goingTo direction, newMin int64) {
	nr, nc := mm.goDirection(r, c, goingTo)

	// update this cell
	mm.m[nr][nc] = newMin

	comingFrom := goingTo.opposite()
	dirs := comingFrom.others()
	for _, currDirection := range dirs {
		if mm.canGo(nr, nc, currDirection) {
			chechRes, newCurrMin := mm.checkNewMin(nr, nc, currDirection)
			if chechRes {
				mm.updateCell(nr, nc, currDirection, newCurrMin)
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
