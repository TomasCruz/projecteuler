package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb067(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8B0lS/zJwAJZbB4JEaIhP6Vuy/4=",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
