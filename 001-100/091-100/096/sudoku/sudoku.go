package sudoku

import (
	"runtime"
	"sync"
)

type coord struct {
	r, c int
}

// Sudoku is a sudoku puzzle struct
type Sudoku struct {
	numSolved  int
	val        [9][9]int
	marks      [9][9]map[int]struct{}
	digitMarks [10]map[coord]struct{}
}

// NewSudoku creates new sudoku puzzle
func NewSudoku(str string) (s Sudoku) {
	for i := 1; i < 10; i++ {
		s.digitMarks[i] = make(map[coord]struct{})
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.marks[i][j] = make(map[int]struct{})

			x := (int)(str[9*i+j] - '0')
			s.val[i][j] = x
			if x == 0 {
				// first, all unsolved cells have all possibilities
				for k := 1; k < 10; k++ {
					s.marks[i][j][k] = struct{}{}
					s.digitMarks[k][coord{r: i, c: j}] = struct{}{}
				}
			}
		}
	}

	// eliminate impossible marks
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.val[i][j] != 0 {
				c := coord{r: i, c: j}
				s.solveCell(c, s.val[i][j])
			}
		}
	}

	return
}

// Solved returns if sudoku had been solved
func (s *Sudoku) Solved() bool {
	return s.numSolved == 81
}

// FirstThree returns three digit number in first row of first block
func (s *Sudoku) FirstThree() int {
	return 100*s.val[0][0] + 10*s.val[0][1] + s.val[0][2]
}

// Solve solves the puzzle
func (s *Sudoku) Solve() {
	// heuristics
	loop := true
	for loop {
		loop = false

		// simple heuristics
		loop = loop || s.uniqueMarksInCell()
		loop = loop || s.singleMarksInRows()
		loop = loop || s.singleMarksInColumns()
		loop = loop || s.singleMarksInBlocks()
	}

	if s.Solved() {
		return
	}

	// sub-puzzles
	var subPuzzles []*Sudoku
	uns := s.firstUnsolved()
	wg := sync.WaitGroup{}
	threadPerCoreGuard := make(chan int, runtime.GOMAXPROCS(runtime.NumCPU()))

	for x := range s.marks[uns.r][uns.c] {
		s1 := &Sudoku{}
		s1.copy(s)
		s1.solveCell(uns, x)
		subPuzzles = append(subPuzzles, s1)

		threadPerCoreGuard <- 1
		wg.Add(1)
		go func(sud *Sudoku) {
			sud.Solve()
			<-threadPerCoreGuard
			wg.Done()
		}(s1)
	}
	wg.Wait()

	for _, sx := range subPuzzles {
		if sx.Solved() {
			s.copy(sx)
			return
		}
	}
}

func (s *Sudoku) firstUnsolved() (c coord) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.val[i][j] != 0 {
				continue
			}

			c.r = i
			c.c = j
			return
		}
	}

	return
}

func (s *Sudoku) copy(sSource *Sudoku) {
	s.numSolved = sSource.numSolved

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.val[i][j] = sSource.val[i][j]
			s.marks[i][j] = make(map[int]struct{})
			for x := range sSource.marks[i][j] {
				s.marks[i][j][x] = struct{}{}
			}
		}

		s.digitMarks[i+1] = make(map[coord]struct{})
		for c := range sSource.digitMarks[i+1] {
			s.digitMarks[i+1][c] = struct{}{}
		}
	}
}
