package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb102(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
