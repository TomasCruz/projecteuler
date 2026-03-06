package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb115(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hkqp5vcpX0mfTh+hwr3oDLuaA==",
		calc,
		50,
	); err != nil {
		t.Errorf(err.Error())
	}
}
