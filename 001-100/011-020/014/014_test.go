package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb014(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/xwlT4LjTiG8Z1+ZZ68p4Gd9NXp2rg==",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
