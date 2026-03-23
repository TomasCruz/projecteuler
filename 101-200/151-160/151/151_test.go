package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb151(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9wEmTo/pHq7pi0Nuz4lULtu2JZ1hg7Xh",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
