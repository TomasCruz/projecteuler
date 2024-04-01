package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb081(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8x0lS4jtWVs5uwvrkiS1Er3YxxS1bg==",
		calc,
		80,
	); err != nil {
		t.Errorf(err.Error())
	}
}
