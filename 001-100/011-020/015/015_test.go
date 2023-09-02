package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb015(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hwlQI/sEqX4SdPQ+LfTPxnWjhhO8rfgsYKY9A==",
		calc,
		20,
	); err != nil {
		t.Errorf(err.Error())
	}
}
