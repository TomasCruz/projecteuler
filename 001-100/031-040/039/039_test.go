package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb039(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoD5KOE4SloQbIQJrl5oQPRQCMGQ==", calc, 1000); err != nil {
		t.Errorf(err.Error())
	}
}
