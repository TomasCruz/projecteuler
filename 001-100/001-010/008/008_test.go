package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb008(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RwnSY/sFaPwQdG1tuV/n7lEukDoif4CNTqV",
		calc,
		13,
	); err != nil {
		t.Errorf(err.Error())
	}
}
