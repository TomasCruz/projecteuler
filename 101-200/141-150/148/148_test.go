package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb148(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9R4gQYLtF6H1RNLRyrJEXAgohgxaFj1HsvGQqoLtvag=",
		calc,
		9,
	); err != nil {
		t.Errorf(err.Error())
	}
}
