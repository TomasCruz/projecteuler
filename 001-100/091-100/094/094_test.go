package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb094(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8h4qTIviFKP2pBnS7VtyIsomh90aOqOzPA==",
		calc,
		1000000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
