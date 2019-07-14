package projecteuler

import (
	"fmt"
	"time"
)

// Timed executes the function and displays its execution time
func Timed(f func(...interface{}) error, args ...interface{}) {
	start := time.Now()
	if err := f(args...); err == nil {
		fmt.Println("Execution lasted: ", time.Since(start))
	}
}
