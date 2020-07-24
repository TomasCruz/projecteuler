package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb061(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZ6IeNuxwW2imXXdDVRgPurZ/oX7",
		calc, 6, projecteuler.ConcurrentKind); err != nil {

		t.Errorf(err.Error())
	}
}
