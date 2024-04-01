package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb082(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RkiS4nu6iAUY9ZZmVKC9BSDwUeCdQ==",
		calc,
		80,
	); err != nil {
		t.Errorf(err.Error())
	}
}
