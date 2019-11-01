package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb008(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZWLcdtcLHBQB4DpzLd5zVyU7YhUnnGzpLOg", calc, 13); err != nil {
		t.Errorf(err.Error())
	}
}
