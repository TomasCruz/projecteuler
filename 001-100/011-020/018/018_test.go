package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb018(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h8lTI5EEeR1g40YxWv59gFBEgw=",
		calc,
		15,
	); err != nil {
		t.Errorf(err.Error())
	}
}
