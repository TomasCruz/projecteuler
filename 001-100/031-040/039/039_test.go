package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb039(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xsiRdp1BkH8kHQY1eZP2jx0Qg==",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
