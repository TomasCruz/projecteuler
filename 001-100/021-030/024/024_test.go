package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb024(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RgqS4LrEqP2QeFqZcJ15ZSfSh+eELVywl4=",
		calc,
		byte(10),
		1000000,
	); err != nil {

		t.Errorf(err.Error())
	}
}
