package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb086(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hcjQOhQshKIkKleHY8X7mqs920=",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
