package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb122(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hoqSt5CxPpURGVP98cs6FMlPoQ=",
		calc,
		200,
	); err != nil {
		t.Errorf(err.Error())
	}
}
