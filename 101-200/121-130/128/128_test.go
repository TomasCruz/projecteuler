package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb128(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hsnSY3iFaPyQ9FapQHTllT4AdHpPxfFlhip",
		calc,
		2000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
