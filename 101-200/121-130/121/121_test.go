package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb121(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9R0kQWMXRvWE29D18g8z8XfHntI=",
		calc,
		15,
	); err != nil {
		t.Errorf(err.Error())
	}
}
