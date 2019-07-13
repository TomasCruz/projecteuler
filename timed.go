package projecteuler

import (
	"fmt"
	"time"
)

// executes the function and displays its execution time
func Timed(f func(...interface{}), args ...interface{}) {
	start := time.Now()
	f(args...)
	fmt.Println("Execution lasted: ", time.Since(start))
}
