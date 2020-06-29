package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb054(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBJGIfSs7SJ6QPMCa5g2vVytaAw==", calc, "p054_poker.txt"); err != nil {
		t.Errorf(err.Error())
	}
}
