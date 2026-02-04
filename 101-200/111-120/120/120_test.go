package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb120(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9BwhSIPoEqfwrvRTA1TZVPfD/CvVOdtE2Q==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
