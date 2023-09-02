package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb030(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8xshQIjjGNcUmsIRqRWLTVDtYsr5CQ==",
		calc,
		5,
	); err != nil {
		t.Errorf(err.Error())
	}
}
