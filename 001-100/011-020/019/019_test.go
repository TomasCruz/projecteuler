package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb019(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9hgj7hdn1agx2VcxtNL3QmDsZw==",
		calc,
		100,
	); err != nil {
		t.Errorf(err.Error())
	}
}
