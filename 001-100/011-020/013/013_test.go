package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb013(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8hohT4jtEaXzQUn+mB6zdLKlJGkQui3sm8k=",
		calc,
		10,
	); err != nil {
		t.Errorf(err.Error())
	}
}
