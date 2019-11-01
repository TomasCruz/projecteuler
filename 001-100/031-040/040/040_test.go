package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb040(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZeOjWp18jsos/n0/KgTP95edA==", calc, 6); err != nil {
		t.Errorf(err.Error())
	}
}
