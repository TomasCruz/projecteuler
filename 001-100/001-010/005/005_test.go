package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb005(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RwgT4LoEqHw24dRNMV/llPriqasOInSbQ==",
		calc,
		20,
	); err != nil {
		t.Errorf(err.Error())
	}
}
