package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb098(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hclToI44gFxU3hckI58F7rDVWWc",
		calc,
		10,
	); err != nil {
		t.Errorf(err.Error())
	}
}
