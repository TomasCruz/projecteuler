package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb076(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hYiTY3jFa7xB7MrD4Ol8HfjjDBWaZrwxA==",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
