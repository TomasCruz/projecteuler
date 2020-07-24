package main

type cyclicalFigurate struct {
	wantedSetSize int
	s             figurateMap
	successors    [][]figurateNumber
}

func makeCyclicalFigurate(setSize int) cyclicalFigurate {
	cf := cyclicalFigurate{
		wantedSetSize: setSize,
		s:             make(figurateMap),
		successors:    make([][]figurateNumber, 100),
	}

	cf.generateFigurateMap()
	cf.generateSuccessors()

	return cf
}

func appendSet(lhs, rhs figurateMap) {
	for x, s := range rhs {
		if f, ok := lhs[x]; ok {
			s.t |= f.t
			lhs[x] = s
		} else {
			lhs[x] = s
		}
	}
}

func (cf *cyclicalFigurate) generateFigurateMap() {
	fs := generateTriangles()
	appendSet(fs, generateSquares())
	appendSet(fs, generatePentagonal())

	if cf.wantedSetSize != 3 {
		appendSet(fs, generateHexagonal())
		appendSet(fs, generateHeptagonal())
		appendSet(fs, generateOctagonal())
	}

	cf.s = fs
	return
}

func (cf *cyclicalFigurate) generateSuccessors() {
	for x, f := range cf.s {
		cf.successors[x/100] = append(cf.successors[x/100], f)
	}
}

func (cf cyclicalFigurate) jobFunction(f figurateNumber) (sum int, err error) {
	s := make([]figurateNumber, 0, cf.wantedSetSize)
	s = append(s, f)
	sum = cf.checkSlice(s)
	return
}

func (cf cyclicalFigurate) checkSlice(s []figurateNumber) (sum int) {
	rSoFar := cf.represented(s[:len(s)-1])
	if rSoFar == cf.represented(s) {
		return
	}

	if len(s) == cf.wantedSetSize-1 {
		lastInt := 100*(s[len(s)-1].x%100) + s[0].x/100
		last, ok := cf.s[lastInt]
		if !ok {
			return
		}

		rSoFar = cf.represented(s)
		s = append(s, last)
		if rSoFar != cf.represented(s) {
			sum = cf.checkIfSolution(s)
		}
		s = s[:len(s)-1]

		return
	}

	suc := cf.successors[s[len(s)-1].x%100]
	for _, currSuc := range suc {
		s = append(s, currSuc)
		sum = cf.checkSlice(s)
		s = s[:len(s)-1]

		if sum != 0 {
			return
		}
	}

	return
}

func (cf cyclicalFigurate) checkIfSolution(s []figurateNumber) int {
	target := 1<<(cf.wantedSetSize) - 1
	if cf.represented(s) != target {
		return 0
	}

	result := 0
	for _, f := range s {
		result += f.x
	}

	return result
}

func (cf cyclicalFigurate) represented(s []figurateNumber) int {
	result := 0

	for _, f := range s {
		result |= f.t
	}

	return result
}
