package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb152(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9B8jkZMBflhtxiF/oWMPcHgiNA==",
		calc,
		80,
	); err != nil {
		t.Error(err.Error())
	}
}
