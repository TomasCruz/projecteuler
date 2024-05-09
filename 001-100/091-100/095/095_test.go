package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb095(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hshSY3jb1F0pCuCCYE2/nb7Nfu+",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
