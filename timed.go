package projecteuler

import (
	"fmt"
	"time"
)

func printCalc(f func(...interface{}) (string, error), args ...interface{}) (err error) {
	var result string

	if result, err = f(args...); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	return
}

// Timed executes the function and displays its execution time
func Timed(f func(...interface{}) (string, error), args ...interface{}) {
	start := time.Now()
	if err := printCalc(f, args...); err == nil {
		fmt.Println("Execution lasted: ", time.Since(start))
	}
}
