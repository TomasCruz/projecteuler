package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb058(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RkgTIpfXHSD+gHHqNxumMu9CQ7q",
		calc,
		15000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
