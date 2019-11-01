package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb018(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpaJdGWX4lfsf4fO7fOpRic5svI=", calc, 15); err != nil {
		t.Errorf(err.Error())
	}
}
