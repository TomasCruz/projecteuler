package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb093(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h0nQFYvqkKZzS3LeQoFDUVBKCU=",
		calc,
		10,
	); err != nil {
		t.Errorf(err.Error())
	}
}
