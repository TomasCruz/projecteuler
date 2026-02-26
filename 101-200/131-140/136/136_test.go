package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb136(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RomTI7vHkidoahGFL4EJ05JQ9RyTTE=",
		calc,
		50000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
