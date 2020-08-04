package projecteuler

// ReverseSquares calculates and returns a map of squares to roots for roots less than limit
func ReverseSquares(limit int) (roots map[int]int) {
	roots = make(map[int]int)
	for i := 0; i < limit; i++ {
		roots[i*i] = i
	}

	return
}
