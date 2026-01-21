package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb107(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RorTozjSJNnRFnRqSqq7bIubk8aWg==",
		calc,
		1,
	); err != nil {
		t.Errorf(err.Error())
	}
}
