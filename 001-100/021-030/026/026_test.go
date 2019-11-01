package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb026(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoDp6No0EFy/BNL1Do7TufMY9WvA==", calc, 1000); err != nil {
		t.Errorf(err.Error())
	}
}
