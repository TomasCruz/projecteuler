package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb129(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h8iSIvoFOSDrlw4zMQc91qMM2pun74=",
		calc,
		6,
	); err != nil {
		t.Errorf(err.Error())
	}
}
