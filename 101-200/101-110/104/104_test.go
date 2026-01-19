package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb104(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9B0rTI3iX4phdTL3z7G9aEXBzaW45g==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
