package sudoku

func (s *Sudoku) singleMarksInRows() (madeChanges bool) {
	for r := 0; r < 9; r++ {
		madeChanges = madeChanges || s.singleMarksInSpecificRow(r)
	}

	return
}

func (s *Sudoku) singleMarksInColumns() (madeChanges bool) {
	for c := 0; c < 9; c++ {
		madeChanges = madeChanges || s.singleMarksInSpecificColumn(c)
	}

	return
}

func (s *Sudoku) singleMarksInBlocks() (madeChanges bool) {
	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			madeChanges = madeChanges || s.singleMarksInSpecificBlock(coord{r: r, c: c})
		}
	}

	return
}

func (s *Sudoku) singleMarksInSpecificRow(r int) bool {
	digitMarks := s.getMarksInSpecificRow(r)
	return s.singleMarksInSet(digitMarks)
}

func (s *Sudoku) singleMarksInSpecificColumn(c int) bool {
	digitMarks := s.getMarksInSpecificColumn(c)
	return s.singleMarksInSet(digitMarks)
}

func (s *Sudoku) singleMarksInSpecificBlock(crd coord) bool {
	digitMarks := s.getMarksInSpecificBlock(crd)
	return s.singleMarksInSet(digitMarks)
}

func (s *Sudoku) singleMarksInSet(rowDigitMarks [10]map[coord]struct{}) bool {
	for x := 1; x < 10; x++ {
		if len(rowDigitMarks[x]) == 1 {
			for c := range rowDigitMarks[x] {
				s.solveCell(c, x)
				return true
			}
		}
	}

	return false
}
