package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb040(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9R4iYM70P6QUqL8Biv9Loo8yzw==",
		calc,
		6,
	); err != nil {
		t.Errorf(err.Error())
	}
}
