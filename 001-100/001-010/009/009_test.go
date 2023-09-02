package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb009(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9B4qT47qF6cqV4offw6NBQDbI8QwYaFP",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
