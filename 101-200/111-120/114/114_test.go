package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb114(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hkmT47sE6fwRdhtxE70sgruznXM6KvXU8AC",
		calc,
		50,
	); err != nil {
		t.Errorf(err.Error())
	}
}
