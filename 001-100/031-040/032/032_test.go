package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb032(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoA5OMcte+6QlzRtAF5bUINd89gcwb", calc); err != nil {
		t.Errorf(err.Error())
	}
}
