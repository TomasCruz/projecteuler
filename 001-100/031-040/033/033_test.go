package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb033(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h8iQ7Fc36mO9bRLyXnfg/GrNA==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
