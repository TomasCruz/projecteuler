package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb046(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8hglT6J6BMudS7uS0ZIVp1UATPg=",
		calc,
		10000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
