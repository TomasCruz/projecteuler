package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb090(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h0jT6gBGUaFsnUf6kkFa9ya2Xg=",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
