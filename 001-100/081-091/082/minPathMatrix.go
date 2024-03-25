package main

import "math"

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
	for i := r + 1; i < mm.limit; i++ {
		mm.m[i][0] = mm.m[i-1][0] + mm.input[i][0]
	}

	for i := 0; i < mm.limit; i++ {
		for j := 1; j < mm.limit; j++ {
			mm.m[i][j] = mm.m[i][j-1] + mm.input[i][j]
		}
	}

	for i := 0; i < mm.limit; i++ {
		for j := 0; j < mm.limit; j++ {
			// mm.updateCell(i, j, left)
			mm.updateCell(i, j, down)
			mm.updateCell(i, j, right)
			mm.updateCell(i, j, up)
		}
	}
}

func (mm minPathMatrix) checkNewMin(r, c int, goingTo direction) (bool, int64) {
	if !mm.canGo(r, c, goingTo) {
		return false, 0
	}

	nr, nc := mm.goDirection(r, c, goingTo)
	if mm.m[r][c] == math.MaxInt64 {
		return false, 0
	}

	newMin := mm.m[r][c] + mm.input[nr][nc]
	return newMin < mm.m[nr][nc], newMin
}

func (mm minPathMatrix) updateCell(r, c int, goingTo direction) {
	chechRes, newMin := mm.checkNewMin(r, c, goingTo)
	if !chechRes {
		return
	}

	comingFrom := goingTo.opposite()
	nr, nc := mm.goDirection(r, c, goingTo)

	// update this cell
	mm.m[nr][nc] = newMin

	dirs := comingFrom.others()
	for _, currDirection := range dirs {
		mm.updateCell(nr, nc, currDirection)
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
