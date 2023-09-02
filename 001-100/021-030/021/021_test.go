package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb021(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9B4kSo1CLioO3Y7QjhClZR1oSUqL",
		calc,
		10000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
