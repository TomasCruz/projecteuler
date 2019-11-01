package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb015(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpWJeNtcK3ZYD4Lj3c9IutobxuR8/AQ2pVTAGQ==", calc, 20); err != nil {
		t.Errorf(err.Error())
	}
}
