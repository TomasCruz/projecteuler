package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb105(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo8BwlSIkGLwh2pkyLNVmj9A8agnTw",
		calc,
		1,
	); err != nil {
		t.Errorf(err.Error())
	}
}
