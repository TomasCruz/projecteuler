package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb025(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8xgqSl/O7PL6hi/9tU5df6XKmvs=",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
