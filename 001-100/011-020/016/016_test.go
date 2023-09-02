package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb016(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hwkTiJGblU332eclX6C62E3Rk0=",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
