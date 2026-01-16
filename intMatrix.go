package projecteuler

import (
	"fmt"
)

// PrintMatrix prints matrix
func PrintMatrix(m [][]int64) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j], " ")
		}
		fmt.Println()
	}
}
