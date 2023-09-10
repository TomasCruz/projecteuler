package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb075(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hkjTo3tUIG4GW3MjR2adIA933sZBw==",
		calc,
		1500000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
