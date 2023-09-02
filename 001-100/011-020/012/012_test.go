package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb012(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8BknT43vF6d7xoNcvmkkDmP1lTWIQ5Jx",
		calc,
		500,
	); err != nil {
		t.Errorf(err.Error())
	}
}
