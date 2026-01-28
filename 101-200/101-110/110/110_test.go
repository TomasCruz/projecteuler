package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb110(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/hwnSIrpF6f0SNnWzrBHXpZUGoP0DnGMnkI6Lvku/Wk=",
		calc,
		4000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
