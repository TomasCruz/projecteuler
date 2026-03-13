package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb126(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hcnSok2LRhgNVBYcrdhcT/FDsgk",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
