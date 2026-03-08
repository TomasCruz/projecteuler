package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb137(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h4gSIruHqH1SdbWzjLwR3fCcWz6k5sEpMXN/p8=",
		calc,
		15,
	); err != nil {
		t.Errorf(err.Error())
	}
}
