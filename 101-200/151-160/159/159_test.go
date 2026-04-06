package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb159(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hsmQILrEq7+jr/P0uF9YCR3w/L4m+4y",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
