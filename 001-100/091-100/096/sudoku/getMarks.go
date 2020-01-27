package sudoku

func (s *Sudoku) getMarksInSpecificRow(r int) (digitMarks [10]map[coord]struct{}) {
	coordSet := make(map[coord]struct{})
	s.rowUnsolvedCoords(r, coordSet)
	return s.getMarksInSpecificCoordSet(coordSet)
}

func (s *Sudoku) getMarksInSpecificColumn(c int) (digitMarks [10]map[coord]struct{}) {
	coordSet := make(map[coord]struct{})
	s.columnUnsolvedCoords(c, coordSet)
	return s.getMarksInSpecificCoordSet(coordSet)
}

func (s *Sudoku) getMarksInSpecificBlock(crd coord) (digitMarks [10]map[coord]struct{}) {
	coordSet := make(map[coord]struct{})
	s.blockUnsolvedCoords(crd, coordSet)
	return s.getMarksInSpecificCoordSet(coordSet)
}

func (s *Sudoku) getMarksInSpecificCoordSet(coordSet map[coord]struct{}) (digitMarks [10]map[coord]struct{}) {
	for x := 1; x < 10; x++ {
		digitMarks[x] = make(map[coord]struct{})
		for crd := range s.digitMarks[x] {
			if _, ok := coordSet[crd]; ok {
				digitMarks[x][crd] = struct{}{}
			}
		}
	}

	return
}
