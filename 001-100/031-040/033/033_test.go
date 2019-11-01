package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb033(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpaOqQ7R8LhA4KdTqXNVdIfJTw==", calc); err != nil {
		t.Errorf(err.Error())
	}
}
