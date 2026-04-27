package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb153(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hgrT4roEqPxQ9PTyLZBXTlRhEYZyVz4lFBYLD2Fw+XR",
		calc,
		100000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
