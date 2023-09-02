package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb063(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8xbK2MOjfBipD57jZUy/j4Ta",
		calc,
		5,
	); err != nil {
		t.Errorf(err.Error())
	}
}
