package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb065(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9Rgg0eQLbmQPnAUXHnhGis3P/A==",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
