package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb074(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8x8gZyTPJCblPQJu42frCOzLsg==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
