package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb138(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h4jQIvuHqX5QdXXzb9EXOGVhYKH09r1/fVCJkbnqU4=",
		calc,
		12,
	); err != nil {
		t.Errorf(err.Error())
	}
}
