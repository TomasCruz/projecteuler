package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb057(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hohSw7cJomYW1NWF71UP5Iong==",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
