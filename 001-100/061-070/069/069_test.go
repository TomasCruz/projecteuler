package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb069(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8h4iTYrqiFDvFQutO8gPVi8NlKgfxQ==",
		calc,
		1000001,
	); err != nil {
		t.Errorf(err.Error())
	}
}
