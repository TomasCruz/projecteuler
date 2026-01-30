package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb111(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8R4gTIvtEqH3RtDVvKUlrptFN3kWQz9DWlUtBg==",
		calc,
		10,
	); err != nil {
		t.Errorf(err.Error())
	}
}
