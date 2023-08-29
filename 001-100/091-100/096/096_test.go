package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb096(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RslSIl3nLX0UVDalvZQlx0ssec3",
		calc,
		"p096_sudoku.txt",
	); err != nil {
		t.Errorf(err.Error())
	}
}
