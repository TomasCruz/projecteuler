package sudoku

func (s *Sudoku) uniqueMarksInCell() (madeChanges bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			madeChanges = madeChanges || s.uniqueMarkCell(coord{r: i, c: j})
		}
	}

	return
}

func (s *Sudoku) uniqueMarkCell(c coord) bool {
	if len(s.marks[c.r][c.c]) == 1 {
		for x := range s.marks[c.r][c.c] {
			s.solveCell(c, x)
			return true
		}
	}

	return false
}
