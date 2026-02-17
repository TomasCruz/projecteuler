package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb149(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8h0qTYnrFaOvfEuvBBL59fmwN8+HinQR",
		calc,
		2000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
