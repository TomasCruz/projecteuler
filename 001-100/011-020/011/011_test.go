package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb011(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8B8kSIvsEKMIj9s18kzz8qZge2id546a",
		calc,
		20,
		4,
	); err != nil {
		t.Errorf(err.Error())
	}
}
