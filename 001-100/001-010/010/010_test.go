package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb010(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hsgQYrpH6X4SNPSwKBL1kSg53JKldfWxDelAg==",
		calc,
		2000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
