package projecteuler

// ReverseSquares calculates and returns a map of square roots for numbers less than limit
func ReverseSquares(limit int) (roots map[int]int) {
	roots = make(map[int]int)
	for i := 0; i < limit; i++ {
		roots[i*i] = i
	}

	return
}
