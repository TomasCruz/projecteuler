package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb097(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xghQYLjFaL3RmW0xRBVYAzuQcw7iUy9pZc=",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
