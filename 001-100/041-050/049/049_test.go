package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb049(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RYkQY3oHq75R9PZpY/cQDeub6Itxs0KgIRsYg==",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
