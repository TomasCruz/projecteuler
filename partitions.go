package projecteuler

func GeneratePartitionMatrix(limit int) [][]int {
	// init
	m := make([][]int, limit+1)
	for i := 0; i <= limit; i++ {
		m[i] = make([]int, limit+1)
	}

	m[1][1] = 1
	for i := 2; i <= limit; i++ {
		m[i][1] = 1
		m[i][2] = i / 2
		m[i][i-1] = 1
		m[i][i] = 1
	}

	for i := 5; i <= limit; i++ {
		for j := 3; j < i-1; j++ {
			sum := 0
			for k := 1; k <= j; k++ {
				sum += m[i-j][k]
			}
			m[i][j] = sum
		}
	}

	return m
}
