package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb096(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZKJcN0Y0Pqe3JPkg11pBggDResP", calc, "p096_sudoku.txt"); err != nil {
		t.Errorf(err.Error())
	}
}
