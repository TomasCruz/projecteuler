package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb132(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xshSoLseO+S6cG1chezA3YPVCAamA==",
		calc,
		40,
	); err != nil {
		t.Errorf(err.Error())
	}
}
