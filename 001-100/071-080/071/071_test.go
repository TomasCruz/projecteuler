package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb071(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8x0qTYzqmHprXcCQ9FjxdYR8t6Xm2Q==",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
