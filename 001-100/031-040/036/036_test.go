package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb036(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xggSYPtZF90H94c2W+aeJNsH8oKCw==",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
