package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb065(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZGM3U68uj4uY19JsgOls0jtWg==", calc, 100); err != nil {
		t.Errorf(err.Error())
	}
}
