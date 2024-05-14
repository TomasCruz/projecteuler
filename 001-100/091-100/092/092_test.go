package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb092(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xoqSYruEQ9ekYgVQiCsntmGchX09ts=",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
