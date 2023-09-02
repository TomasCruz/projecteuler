package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb006(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RojTo/rEqeG3BGDJx/jjhxfgrFEkNi/",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
