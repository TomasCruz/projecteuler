package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb113(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8h4jTorqEq/xQtXSy7b5qQOJ8HZj7CE66yqWQI+/",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
