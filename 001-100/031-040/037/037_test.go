package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb037(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8BsqS4rtnSdxPqg+vDzB3AyRGds0ug==",
		calc,
	); err != nil {
		t.Errorf(err.Error())
	}
}
