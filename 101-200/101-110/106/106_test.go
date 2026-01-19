package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb106(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9R4hQI8LEapcCcxDxaPKsVY/eDKp",
		calc,
		12,
	); err != nil {
		t.Errorf(err.Error())
	}
}
