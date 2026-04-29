package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb123(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9R4iS477Q/TchunVW6RPjn3RHbn6",
		calc,
		10,
	); err != nil {
		t.Error(err.Error())
	}
}
