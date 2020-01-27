package sudoku

func (s *Sudoku) rowUnsolvedCoords(r int, coordSet map[coord]struct{}) {
	for c := 0; c < 9; c++ {
		if s.val[r][c] != 0 {
			continue
		}

		coordSet[coord{r: r, c: c}] = struct{}{}
	}
}

func (s *Sudoku) columnUnsolvedCoords(c int, coordSet map[coord]struct{}) {
	for r := 0; r < 9; r++ {
		if s.val[r][c] != 0 {
			continue
		}

		coordSet[coord{r: r, c: c}] = struct{}{}
	}
}

func (s *Sudoku) blockUnsolvedCoords(c coord, coordSet map[coord]struct{}) {
	blockIndex := 3*(c.r/3) + c.c/3
	rowStart := 3 * (blockIndex / 3)
	colStart := 3 * (blockIndex % 3)

	for row := rowStart; row < rowStart+3; row++ {
		for col := colStart; col < colStart+3; col++ {
			if s.val[row][col] != 0 {
				continue
			}

			coordSet[coord{r: row, c: col}] = struct{}{}
		}
	}
}
