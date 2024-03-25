package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb083(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8x0nSYPvXHqFF4sppfwRN8mu5S2T6Q==",
		calc,
		80,
	); err != nil {
		t.Errorf(err.Error())
	}
}
