package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb078(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8hohT49zO7WXUaBRzB4pvJxZbReT",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
