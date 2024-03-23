package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb089(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8Bsh4+T+n1oZF7VmRr4AwtsJ/w==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
