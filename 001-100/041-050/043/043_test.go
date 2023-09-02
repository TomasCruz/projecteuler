package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb043(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hkkQY7pFKP4SNHVSBK1GhFw6B1sFivKYE0N",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
