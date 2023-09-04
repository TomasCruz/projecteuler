package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb070(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xwjQYPoFFtPYOXv5bP1e4W6S7aUZTw=",
		calc,
		10000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
