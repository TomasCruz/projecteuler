package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb066(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8RkjBCxt1BcdOi1MH/o20ye/gQ==",
		calc,
		1000,
		chakravalaSolver{},
	); err != nil {

		t.Errorf(err.Error())
	}
}
