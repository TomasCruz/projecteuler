package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb045(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hohS4ztEa/wRHg8CLDinIqdqdEuV5iYm3A=",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
