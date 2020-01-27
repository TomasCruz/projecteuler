package sudoku

func (s *Sudoku) solveCell(c coord, val int) {
	s.numSolved++
	s.val[c.r][c.c] = val

	// delete all from marks and digitMarks
	s.marks[c.r][c.c] = make(map[int]struct{})
	for k := 1; k < 10; k++ {
		delete(s.digitMarks[k], c)
	}

	// remove this mark
	markSet := make(map[int]struct{})
	markSet[val] = struct{}{}

	// ... from this set
	coordSet := make(map[coord]struct{})
	s.rowUnsolvedCoords(c.r, coordSet)
	s.columnUnsolvedCoords(c.c, coordSet)
	s.blockUnsolvedCoords(c, coordSet)

	s.eliminateMarksInSet(markSet, coordSet)
}
