package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb019(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpGPyq9VhPMAVoULh2yTMSkRkw==", calc, 100); err != nil {
		t.Errorf(err.Error())
	}
}
