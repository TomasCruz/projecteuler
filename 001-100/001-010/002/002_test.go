package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb002(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8xkjS4zpFfGaRatBxprYChxz8T7vtIo=",
		calc,
		int64(4000000),
	); err != nil {
		t.Errorf(err.Error())
	}
}
