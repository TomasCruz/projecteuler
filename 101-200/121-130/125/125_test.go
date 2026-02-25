package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb125(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RYiToLsHqb3SCetysO3RkWPUTtvgDSTCJQ=",
		calc,
		8,
	); err != nil {
		t.Errorf(err.Error())
	}
}
