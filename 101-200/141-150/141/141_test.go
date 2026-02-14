package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb141(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xgqTI7uFKT3QNTZNeSOHXoyMv88FXKH/Bu8NA==",
		calc,
		12,
	); err != nil {
		t.Errorf(err.Error())
	}
}
