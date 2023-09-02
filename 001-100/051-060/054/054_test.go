package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb054(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9BgkDzUQkqpDgS/klWQ7SJYNQg==",
		calc,
		"p054_poker.txt",
	); err != nil {
		t.Errorf(err.Error())
	}
}
