package sudoku

func (s *Sudoku) eliminateMarksInSet(markSet map[int]struct{}, coordSet map[coord]struct{}) {
	for c := range coordSet {
		for x := range markSet {
			delete(s.marks[c.r][c.c], x)
			delete(s.digitMarks[x], c)
		}
	}
}
