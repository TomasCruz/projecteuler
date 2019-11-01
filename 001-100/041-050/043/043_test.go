package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb043(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpCIedpZLXBYDoB/G4w4f92JutU0uJOZkD1A", calc); err != nil {
		t.Errorf(err.Error())
	}
}
