package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb088(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8BoqT4/vECSGk9HlhdRQs8CRcCSWcg0=",
		calc,
		12000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
