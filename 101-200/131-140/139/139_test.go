package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb139(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h8iTYztEaaBtdLgC+SmZ/I6FXjXHoEx",
		calc,
		8,
	); err != nil {
		t.Errorf(err.Error())
	}
}
