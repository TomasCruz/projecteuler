package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb091(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hsgS48i3cq8N2DPDayu1C+/A+pC",
		calc,
		50,
	); err != nil {
		t.Errorf(err.Error())
	}
}
