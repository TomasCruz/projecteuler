package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb150(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo6h0lSYnuH6H4Qd3jOL7x3UMNN3nnhMdy2N0=",
		calc,
		2,
	); err != nil {
		t.Errorf(err.Error())
	}
}
