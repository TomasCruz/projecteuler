package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb077(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8B4lzTMuCos016qCyK+g/WDH",
		calc,
		5000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
