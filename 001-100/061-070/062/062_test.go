package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb062(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h0lSIjvHqL0R9nTpH+VapGjqENJEp730EBaJQ==",
		calc,
		5,
	); err != nil {
		t.Errorf(err.Error())
	}
}
