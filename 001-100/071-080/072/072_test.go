package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb072(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9B8hQY3pEqLyQtjRRZxSzKiUDJz4QzVE7lEK1w==",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
