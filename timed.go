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

// TimedStr executes the function and displays string result and execution time
func TimedStr(f func(...interface{}) (string, error), args ...interface{}) {
	start := time.Now()
	if result, err := f(args...); err == nil {
		fmt.Println(result)
		fmt.Println("Execution lasted: ", time.Since(start))
	}
}
