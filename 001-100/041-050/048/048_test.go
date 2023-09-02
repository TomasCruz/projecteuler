package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb048(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/h4jSIPuEaDwQY0KT/R1y6UZisYeX6oA/Ys=",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
