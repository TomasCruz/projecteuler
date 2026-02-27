package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb143(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9B8lTYPpHqC7/aBbiSJkw9yd1Hv9PsFg",
		calc,
		120000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
