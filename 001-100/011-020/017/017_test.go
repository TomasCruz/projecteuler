package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb017(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9R4jSo/n2rOBc6gpIsrR8FZNWWsN",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
