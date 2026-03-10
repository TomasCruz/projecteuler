package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb140(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8hklS4PpEqT1Q9jZzpTV+vovt3OHi7Fo6YF0ijI=",
		calc,
		30,
	); err != nil {
		t.Errorf(err.Error())
	}
}
