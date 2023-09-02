package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb051(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h0jS4rpCMRFKj/w/fYWoH2CCJ/aVQ==",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
